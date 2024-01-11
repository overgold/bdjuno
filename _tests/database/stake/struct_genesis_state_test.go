package bank

import (
	"encoding/json"
	"fmt"
	"testing"

	"git.ooo.ua/vipcoin/ovg-chain/x/stake/types"
	"github.com/brianvoe/gofakeit/v6"
	sdk "github.com/cosmos/cosmos-sdk/types"

	db "github.com/forbole/bdjuno/v4/_tests/database"
)

const (
	ovgAddress = `ovg1[a-z0-9]{38}$`
)

func TestRepository_StructGenesisState(t *testing.T) {
	type args struct {
		msg []types.GenesisState
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] StructGenesisState",
			args: args{
				msg: []types.GenesisState{
					{
						Params: types.Params{},
						StatsList: []types.Stats{
							{
								Date: "2024-01-10",
								Stats: &types.DailyStats{
									Id:                  0,
									AmountSell:          sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000)),
									AmountBuy:           sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000)),
									AmountIssue:         sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000)),
									CountIssue:          1,
									CountUsersIssue:     1,
									CountSell:           1,
									CountUsersSell:      1,
									CountBuy:            1,
									CountUsersBuy:       1,
									AmountWithdraw:      sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000)),
									CountWithdraw:       1,
									CountUsersWithdraw:  1,
									AmountBurn:          sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000)),
									CountBurn:           1,
									CountUsersBurn:      1,
									CountUsers:          1,
									DistributionsTotal:  sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000)),
									DistributionsByUser: 1,
								},
							},
							{
								Date: "2024-01-11",
								Stats: &types.DailyStats{
									Id:                  1,
									AmountSell:          sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000)),
									AmountBuy:           sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000)),
									AmountIssue:         sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000)),
									CountIssue:          2,
									CountUsersIssue:     2,
									CountSell:           2,
									CountUsersSell:      2,
									CountBuy:            2,
									CountUsersBuy:       2,
									AmountWithdraw:      sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000)),
									CountWithdraw:       2,
									CountUsersWithdraw:  2,
									AmountBurn:          sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000)),
									CountBurn:           2,
									CountUsersBurn:      2,
									CountUsers:          2,
									DistributionsTotal:  sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000)),
									DistributionsByUser: 2,
								},
							},
						},
						DailyStatsCount: 2,
						StakesList: []types.Stakes{
							{
								AccountAddress: gofakeit.Regex(ovgAddress),
								Stake: &types.Stake{
									SellAmount:   sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000)),
									RewardAmount: sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000)),
									StakeAmount:  sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000)),
									Transactions: []*types.Transaction{
										{
											Hash:        gofakeit.LetterN(64),
											AddressFrom: gofakeit.Regex(ovgAddress),
											AddressTo:   gofakeit.Regex(ovgAddress),
											Amount:      sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000)),
											Kind:        types.TRANSFER_KIND_ISSUE,
											Description: gofakeit.EmojiDescription(),
											Timestamp:   "2024-01-10",
										},
										{
											Hash:        gofakeit.LetterN(64),
											AddressFrom: gofakeit.Regex(ovgAddress),
											AddressTo:   gofakeit.Regex(ovgAddress),
											Amount:      sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000)),
											Kind:        types.TRANSFER_KIND_WITHDRAW,
											Description: gofakeit.EmojiDescription(),
											Timestamp:   "2024-01-10",
										},
									},
								},
							},
							{
								AccountAddress: gofakeit.Regex(ovgAddress),
								Stake: &types.Stake{
									SellAmount:   sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000)),
									RewardAmount: sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000)),
									StakeAmount:  sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000)),
									Transactions: []*types.Transaction{
										{
											Hash:        gofakeit.LetterN(64),
											AddressFrom: gofakeit.Regex(ovgAddress),
											AddressTo:   gofakeit.Regex(ovgAddress),
											Amount:      sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000)),
											Kind:        types.TRANSFER_KIND_ISSUE,
											Description: gofakeit.EmojiDescription(),
											Timestamp:   "2024-01-11",
										},
										{
											Hash:        gofakeit.LetterN(64),
											AddressFrom: gofakeit.Regex(ovgAddress),
											AddressTo:   gofakeit.Regex(ovgAddress),
											Amount:      sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000)),
											Kind:        types.TRANSFER_KIND_WITHDRAW,
											Description: gofakeit.EmojiDescription(),
											Timestamp:   "2024-01-11",
										},
									},
								},
							},
						},
						BuyRequestsList: []types.BuyRequests{
							{
								AccountAddress: gofakeit.Regex(ovgAddress),
								BuyRequest:     &types.BuyRequest{Amount: sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000))},
							},
							{
								AccountAddress: gofakeit.Regex(ovgAddress),
								BuyRequest:     &types.BuyRequest{Amount: sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000))},
							},
						},
						SellRequestsList: []types.SellRequests{
							{
								AccountAddress: gofakeit.Regex(ovgAddress),
								SellRequest:    &types.SellRequest{Amount: sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000))},
							},
							{
								AccountAddress: gofakeit.Regex(ovgAddress),
								SellRequest:    &types.SellRequest{Amount: sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000))},
							},
						},
						LimitsList: []types.Limits{
							{
								Denom: "stovg",
								MinAmounts: &types.MinAmounts{
									Sell: sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000)),
									Buy:  sdk.NewCoin("stovg", sdk.NewInt(1_0000_0000)),
								},
								Creator: db.TestAddressCreator,
							},
							{
								Denom: "stovg",
								MinAmounts: &types.MinAmounts{
									Sell: sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000)),
									Buy:  sdk.NewCoin("stovg", sdk.NewInt(2_0000_0000)),
								},
								Creator: db.TestAddressCreator,
							},
						},
						UsersCount: 2,
						UserStatsList: []types.UserStats{
							{
								Date: "2024-01-10",
								UniqueUsers: &types.UniqueUsers{
									Issue:    []string{gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress)},
									Burn:     []string{gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress)},
									Withdraw: []string{gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress)},
									Sell:     []string{gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress)},
									Buy:      []string{gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress)},
									Total:    []string{gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress)},
								},
							},
							{
								Date: "2024-01-11",
								UniqueUsers: &types.UniqueUsers{
									Issue:    []string{gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress)},
									Burn:     []string{gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress)},
									Withdraw: []string{gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress)},
									Sell:     []string{gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress)},
									Buy:      []string{gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress)},
									Total:    []string{gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress), gofakeit.Regex(ovgAddress)},
								},
							},
						},
						TotalFree: 2,
						TotalSell: 2,
						SellParamsList: []types.SellParams{
							{
								Denom: "stovg",
								SellStakeParams: &types.SellStakeParams{
									MinSellRequests: 0,
									MaxSellRequests: 0,
								},
								Creator: db.TestAddressCreator,
							},
							{
								Denom: "stovg",
								SellStakeParams: &types.SellStakeParams{
									MinSellRequests: 1_0000_0000,
									MaxSellRequests: 1_0000_0000,
								},
								Creator: db.TestAddressCreator,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, err := json.MarshalIndent(tt.args.msg, "", "\t")
			if err != nil {
				t.Error(err)
				return
			}

			fmt.Println(string(jsonData))
		})
	}
}
