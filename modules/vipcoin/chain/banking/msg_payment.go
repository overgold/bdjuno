package banking

import (
	"math"
	"strings"
	"time"

	accounts "git.ooo.ua/vipcoin/chain/x/accounts/types"
	assets "git.ooo.ua/vipcoin/chain/x/assets/types"
	banking "git.ooo.ua/vipcoin/chain/x/banking/types"
	wallets "git.ooo.ua/vipcoin/chain/x/wallets/types"
	"git.ooo.ua/vipcoin/lib/filter"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v2/types"

	dbtypes "github.com/forbole/bdjuno/v2/database/types"
)

// handleMsgPayments allows to properly handle a MsgSetState
func (m *Module) handleMsgPayments(tx *juno.Tx, _ int, msg *banking.MsgPayment) error {
	msg.WalletFrom = strings.ToLower(msg.WalletFrom)
	msg.WalletTo = strings.ToLower(msg.WalletTo)
	msg.Asset = strings.ToLower(msg.Asset)

	if err := m.bankingRepo.SaveMsgPayments(msg); err != nil {
		return err
	}

	asset, err := m.assetRepo.GetAssets(filter.NewFilter().SetArgument(dbtypes.FieldName, msg.Asset))
	switch {
	case err != nil:
		return err
	case len(asset) != 1:
		return assets.ErrNotFoundAsset
	}

	walletFrom, err := m.walletsRepo.GetWallets(filter.NewFilter().SetArgument(dbtypes.FieldAddress, msg.WalletFrom))
	switch {
	case err != nil:
		return err
	case len(walletFrom) != 1:
		return wallets.ErrInvalidAddressField
	}

	walletTo, err := m.walletsRepo.GetWallets(filter.NewFilter().SetArgument(dbtypes.FieldAddress, msg.WalletTo))
	switch {
	case err != nil:
		return err
	case len(walletTo) != 1:
		return wallets.ErrInvalidAddressField
	}

	accountFrom, err := m.accountsRepo.GetAccounts(filter.NewFilter().SetArgument(dbtypes.FieldAddress, walletFrom[0].AccountAddress))
	switch {
	case err != nil:
		return err
	case len(accountFrom) != 1:
		return accounts.ErrInvalidAddressField
	}

	accountTo, err := m.accountsRepo.GetAccounts(filter.NewFilter().SetArgument(dbtypes.FieldAddress, walletTo[0].AccountAddress))
	switch {
	case err != nil:
		return err
	case len(accountTo) != 1:
		return accounts.ErrInvalidAddressField
	}

	switch m.getPaymentFeeFromTx(tx, msg) {
	case 0:
		return m.payment(tx, msg, *asset[0], *walletFrom[0], *walletTo[0])
	default:
		return m.paymentWithFee(tx, accountTo[0], msg, *asset[0], *walletFrom[0], *walletTo[0])
	}
}

// payment - creates payment without fee
func (m *Module) payment(
	tx *juno.Tx,
	msg *banking.MsgPayment,
	asset assets.Asset,
	walletFrom, walletTo wallets.Wallet,
) error {
	coin := sdk.NewCoin(asset.Name, sdk.NewIntFromUint64(msg.Amount))

	// subtract coins from sender wallet balance
	walletFrom.Balance = walletFrom.Balance.Sub(sdk.NewCoins(coin))
	if err := m.walletsRepo.UpdateWallets(&walletFrom); err != nil {
		return err
	}

	// add coins to receiver wallet balance
	walletTo.Balance = walletTo.Balance.Add(coin)
	if err := m.walletsRepo.UpdateWallets(&walletTo); err != nil {
		return err
	}

	timestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return err
	}

	payment := &banking.Payment{
		WalletFrom: msg.WalletFrom,
		WalletTo:   msg.WalletTo,
		Fee:        0,
		BaseTransfer: banking.BaseTransfer{
			Asset:     msg.Asset,
			Amount:    msg.Amount,
			Extras:    msg.Extras,
			Kind:      banking.TRANSFER_KIND_PAYMENT,
			Timestamp: timestamp.Unix(),
			TxHash:    tx.TxHash,
		},
	}

	if walletFrom.Kind == wallets.WALLET_KIND_SYSTEM_DEFERRED {
		payment.Kind = banking.TRANSFER_KIND_DEFERRED
	}

	return m.bankingRepo.SavePayments(payment)
}

// paymentWithFee - creates payment with fee
func (m *Module) paymentWithFee(
	tx *juno.Tx,
	to *accounts.Account,
	msg *banking.MsgPayment,
	asset assets.Asset,
	walletFrom, walletTo wallets.Wallet,
) error {
	AssetPolicyShareholderRefReward := asset.CheckPolicy(assets.ASSET_POLICY_SHAREHOLDER_REF_REWARD)

	// Getting supplementary wallets
	walletsSystemReward, err := m.walletsRepo.GetWallets(filter.NewFilter().SetArgument(dbtypes.FieldKind, wallets.WALLET_KIND_SYSTEM_REWARD))
	switch {
	case err != nil:
		return err
	case len(walletsSystemReward) != 1:
		return wallets.ErrInvalidKindField
	}

	walletsRefReward, err := m.walletsRepo.GetWallets(
		filter.NewFilter().
			SetCondition(filter.ConditionAND).
			SetArgument(dbtypes.FieldAccountAddress, walletsSystemReward[0].AccountAddress).
			SetArgument(dbtypes.FieldKind, wallets.WALLET_KIND_REFERRER_REWARD))
	switch {
	case err != nil:
		return err
	case len(walletsRefReward) != 1:
		return wallets.ErrInvalidKindField
	}

	walletsVoid, err := m.walletsRepo.GetWallets(
		filter.NewFilter().
			SetCondition(filter.ConditionAND).
			SetArgument(dbtypes.FieldAccountAddress, walletsSystemReward[0].AccountAddress).
			SetArgument(dbtypes.FieldKind, wallets.WALLET_KIND_VOID))
	switch {
	case err != nil:
		return err
	case len(walletsVoid) != 1:
		return wallets.ErrInvalidKindField
	}

	walletSystemReward := walletsSystemReward[0]
	walletRefReward := walletsRefReward[0]
	walletVoid := walletsVoid[0]

	// ----- Getting referrer for payment receiver and referrer ref reward wallet -----
	// Looking for receiver referrer
	var referrerAccAddr string
	for _, a := range to.Affiliates {
		if a.Affiliation == accounts.AFFILIATION_KIND_REFERRER {
			referrerAccAddr = a.Address
			break
		}
	}

	// If referrer exists then checking the shareholder state
	// And setting user`s ref reward wallet instead of system ref reward wallet
	if referrerAccAddr != "" {
		// get wallet owner
		referrerAccounts, err := m.accountsRepo.GetAccounts(filter.NewFilter().SetArgument(dbtypes.FieldAddress, referrerAccAddr))
		switch {
		case err != nil:
			return err
		case len(referrerAccounts) != 1:
			return accounts.ErrInvalidAddressField
		}

		referrerAcc := referrerAccounts[0]

		reqWalletRefReward := wallets.Wallet{
			AccountAddress: referrerAccAddr,
			Kind:           wallets.WALLET_KIND_REFERRER_REWARD,
		}

		if referrerAcc.State == accounts.ACCOUNT_STATE_ACTIVE {
			// check referrer`s account extra if shareholder field is set
			// get stored extra
			if accounts.IsKind(accounts.ACCOUNT_KIND_SHAREHOLDER, referrerAcc.Kinds...) || !AssetPolicyShareholderRefReward {
				// iterate over account`s wallets and look for ref reward wallet
				for _, walletAddr := range referrerAcc.Wallets {
					w, err := m.walletsRepo.GetWallets(filter.NewFilter().SetArgument(dbtypes.FieldAddress, walletAddr))
					switch {
					case err != nil:
						return err
					case len(w) != 1:
						return wallets.ErrInvalidAddressField
					}
					if w[0].OverlapBy(reqWalletRefReward) {
						// Setting user`s ref reward wallet instead of system ref reward wallet
						walletRefReward = w[0]
						break
					}
				}
			}
		}
	}

	// ----- General payment -----
	var (
		feeRaw          = float64(msg.Amount) / 100.0 * (float64(asset.Properties.FeePercent) / 100.0) // FeePercent 100 = 1%
		feeSysRewardRaw = feeRaw / 100.0 * 50.0
		feeRefRewardRaw = feeRaw / 100.0 * 25.0
	)

	var (
		fee          = uint64(math.Round(feeRaw))
		feeSysReward = uint64(math.Round(feeSysRewardRaw))
		feeRefReward = uint64(math.Round(feeRefRewardRaw))
		feeVoid      = fee - (feeSysReward + feeRefReward)
	)

	var (
		coin             = sdk.NewCoin(asset.Name, sdk.NewIntFromUint64(msg.Amount))
		coinAfterFee     = sdk.NewCoin(asset.Name, sdk.NewIntFromUint64(msg.Amount-fee))
		coinFeeSysReward = sdk.NewCoin(asset.Name, sdk.NewIntFromUint64(feeSysReward))
		coinFeeRefReward = sdk.NewCoin(asset.Name, sdk.NewIntFromUint64(feeRefReward))
		coinFeeVoid      = sdk.NewCoin(asset.Name, sdk.NewIntFromUint64(feeVoid))
	)

	// subtract coins from sender wallet balance
	walletFrom.Balance = walletFrom.Balance.Sub(sdk.NewCoins(coin))
	if err := m.walletsRepo.UpdateWallets(&walletFrom); err != nil {
		return err
	}

	// add coins to receiver wallet balance
	walletTo.Balance = walletTo.Balance.Add(coinAfterFee)
	if err := m.walletsRepo.UpdateWallets(&walletTo); err != nil {
		return err
	}

	// add coins to system reward wallet balance
	walletSystemReward.Balance = walletSystemReward.Balance.Add(coinFeeSysReward)
	if err := m.walletsRepo.UpdateWallets(walletSystemReward); err != nil {
		return err
	}

	timestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err != nil {
		return err
	}

	// store sys-reward transfer
	systemReward := &banking.SystemTransfer{
		WalletFrom: msg.WalletTo,
		WalletTo:   walletSystemReward.Address,
		BaseTransfer: banking.BaseTransfer{
			Asset:     msg.Asset,
			Amount:    feeSysReward,
			Kind:      banking.TRANSFER_KIND_SYSTEM_REWARD,
			Timestamp: timestamp.Unix(),
			TxHash:    tx.TxHash,
		},
	}

	if err := m.bankingRepo.SaveSystemTransfers(systemReward); err != nil {
		return err
	}

	// add coins to referrer (if referrer is empty then it will be system ref reward wallet) wallet balance
	walletRefReward.Balance = walletRefReward.Balance.Add(coinFeeRefReward)
	if err := m.walletsRepo.UpdateWallets(walletRefReward); err != nil {
		return err
	}

	// store ref-reward transfer
	systemRefReward := &banking.SystemTransfer{
		WalletFrom: msg.WalletTo,
		WalletTo:   walletRefReward.Address,
		BaseTransfer: banking.BaseTransfer{
			Asset:     msg.Asset,
			Amount:    feeRefReward,
			Kind:      banking.TRANSFER_KIND_SYSTEM_REF_REWARD,
			Timestamp: timestamp.Unix(),
			TxHash:    tx.TxHash,
		},
	}

	if err := m.bankingRepo.SaveSystemTransfers(systemRefReward); err != nil {
		return err
	}

	// add coins to void wallet balance
	walletVoid.Balance = walletVoid.Balance.Add(coinFeeVoid)
	if err := m.walletsRepo.UpdateWallets(walletVoid); err != nil {
		return err
	}

	asset.Burned += feeVoid
	asset.InCirculation -= feeVoid
	if err := m.assetRepo.UpdateAssets(&asset); err != nil {
		return err
	}

	payment := &banking.Payment{
		WalletFrom: msg.WalletFrom,
		WalletTo:   msg.WalletTo,
		Fee:        fee,
		BaseTransfer: banking.BaseTransfer{
			Asset:     msg.Asset,
			Amount:    msg.Amount,
			Extras:    msg.Extras,
			Kind:      banking.TRANSFER_KIND_PAYMENT,
			Timestamp: timestamp.Unix(),
			TxHash:    tx.TxHash,
		},
	}

	if walletFrom.Kind == wallets.WALLET_KIND_SYSTEM_DEFERRED {
		payment.Kind = banking.TRANSFER_KIND_DEFERRED
	}

	return m.bankingRepo.SavePayments(payment)
}
