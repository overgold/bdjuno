package banking

import (
	"encoding/json"
	"errors"
	"strconv"

	banking "git.ooo.ua/vipcoin/chain/x/banking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v2/types"
)

// getPaymentFromTx - get payment from tx
func getPaymentFromTx(tx *juno.Tx, msg *banking.MsgPayment) (*banking.Payment, error) {
	for _, log := range tx.TxResponse.Logs {
		for _, event := range log.Events {
			if event.Type != "vipcoin.chain.banking.Payment" {
				continue
			}

			payment, err := getPaymentTransferFromAttribute(event.Attributes)
			if err != nil {
				return nil, err
			}

			if payment.WalletFrom != msg.WalletFrom {
				continue
			}

			if payment.WalletTo != msg.WalletTo {
				continue
			}

			if payment.BaseTransfer.Amount != msg.Amount {
				continue
			}

			payment.Extras = msg.Extras

			return payment, nil
		}
	}

	return nil, errors.New("not found")
}

// getSystemTransferByKind - get system transfer by kind from logs
func getSystemTransferByKind(tx *juno.Tx, walletFrom string, kind banking.TransferKind) (*banking.SystemTransfer, error) {
	for _, log := range tx.TxResponse.Logs {
		for _, event := range log.Events {
			if event.Type != "vipcoin.chain.banking.SystemTransfer" {
				continue
			}

			transfer, err := getSystemTransferFromAttribute(event.Attributes, kind)
			if err != nil {
				return nil, err
			}

			if transfer.WalletFrom != walletFrom {
				continue
			}

			return transfer, nil
		}

	}

	return nil, errors.New("not found")
}

// getSystemTransferFromAttribute - get system transfer by kind from attributes
func getSystemTransferFromAttribute(attributes []sdk.Attribute, kind banking.TransferKind) (*banking.SystemTransfer, error) {
	for index, attribute := range attributes {
		if attribute.Key != "base_transfer" {
			continue
		}

		var base baseTransfer
		if err := json.Unmarshal([]byte(attribute.Value), &base); err != nil {
			return nil, err
		}

		if base.Kind != kind.String() {
			continue
		}

		if len(attributes) <= (index + 2) {
			return nil, errors.New("invalid attributes")
		}

		if attributes[index+1].Key != "wallet_from" || attributes[index+2].Key != "wallet_to" {
			return nil, errors.New("invalid attributes")
		}

		var walletFrom string
		if err := json.Unmarshal([]byte(attributes[index+1].Value), &walletFrom); err != nil {
			return nil, err
		}

		var walletTo string
		if err := json.Unmarshal([]byte(attributes[index+2].Value), &walletTo); err != nil {
			return nil, err
		}

		baseVipcoin, err := base.toVipcoinBaseTransfer()
		if err != nil {
			return nil, err
		}

		systemTransfer := banking.SystemTransfer{
			BaseTransfer: baseVipcoin,
			WalletFrom:   walletFrom,
			WalletTo:     walletTo,
		}

		return &systemTransfer, nil
	}

	return nil, errors.New("not found")
}

// getPaymentTransferFromAttribute - get payment transfer by kind from attributes
func getPaymentTransferFromAttribute(attributes []sdk.Attribute) (*banking.Payment, error) {
	var base baseTransfer
	if err := json.Unmarshal([]byte(getAttributeValueWithKey(attributes, "base_transfer")), &base); err != nil {
		return nil, err
	}

	var walletFrom string
	if err := json.Unmarshal([]byte(getAttributeValueWithKey(attributes, "wallet_from")), &walletFrom); err != nil {
		return nil, err
	}

	var walletTo string
	if err := json.Unmarshal([]byte(getAttributeValueWithKey(attributes, "wallet_to")), &walletTo); err != nil {
		return nil, err
	}

	var fee string
	if err := json.Unmarshal([]byte(getAttributeValueWithKey(attributes, "fee")), &fee); err != nil {
		return nil, err
	}

	feeUint, err := strconv.ParseUint(fee, 10, 64)
	if err != nil {
		return nil, err
	}

	baseVipcoin, err := base.toVipcoinBaseTransfer()
	if err != nil {
		return nil, err
	}

	paymentTransfer := banking.Payment{
		BaseTransfer: baseVipcoin,
		WalletFrom:   walletFrom,
		WalletTo:     walletTo,
		Fee:          feeUint,
	}

	return &paymentTransfer, nil
}

// getAttributeValueWithKey - get attribute value with key
func getAttributeValueWithKey(attributes []sdk.Attribute, key string) string {
	for _, attribute := range attributes {
		if attribute.Key == key {
			return attribute.Value
		}
	}

	return ""
}
