package banking

import (
	"strconv"
	"strings"

	banking "git.ooo.ua/vipcoin/chain/x/banking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v2/types"
)

// getAttributeValueWithKey - get attribute value with key
func getAttributeValueWithKey(attributes []sdk.Attribute, key string) string {
	for _, attribute := range attributes {
		if attribute.Key == key {
			return attribute.Value
		}
	}

	return ""
}

// getPaymentFeeFromTx - get payment fee from tx
func (m *Module) getPaymentFeeFromTx(tx *juno.Tx, msg *banking.MsgPayment) uint64 {
	var fee string
	for _, log := range tx.TxResponse.Logs {
		for _, event := range log.Events {
			if event.Type != "vipcoin.chain.banking.Payment" {
				continue
			}

			fee = getAttributeValueWithKey(event.Attributes, "fee")
			if fee == "" {
				continue
			}

			if !strings.Contains(getAttributeValueWithKey(event.Attributes, "wallet_from"), msg.WalletFrom) {
				fee = ""
				continue
			}

			if !strings.Contains(getAttributeValueWithKey(event.Attributes, "wallet_to"), msg.WalletTo) {
				fee = ""
				continue
			}

			break
		}
	}

	feeList := m.numberOnly.FindAllString(fee, -1)
	if len(feeList) == 0 {
		return 0
	}

	feeUint, err := strconv.ParseUint(feeList[0], 10, 64)
	if err != nil {
		return 0
	}

	return feeUint
}
