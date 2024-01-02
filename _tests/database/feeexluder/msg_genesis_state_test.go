package feeexluder

import (
	"testing"

	fe "git.ooo.ua/vipcoin/ovg-chain/x/feeexcluder/types"
	"github.com/cosmos/cosmos-sdk/types"

	d "github.com/forbole/bdjuno/v4/_tests/database"
)

func TestRepository_InsertToGenesisState(t *testing.T) {
	type args struct {
		msg []fe.GenesisState
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] InsertToGenesisState",
			args: args{
				msg: []fe.GenesisState{
					{
						Params: fe.Params{},
						AddressList: []fe.Address{
							{
								Id:      1,
								Address: d.TestAddressCreator,
								Creator: d.TestAddressCreator,
							},
							{
								Id:      2,
								Address: "ovg1wvuy80m54dl8qw63u3jnaqjc3y82gnlk36gkjj",
								Creator: d.TestAddressCreator,
							},
						},
						AddressCount: 2,
						DailyStatsList: []fe.DailyStats{
							{
								Id:            1,
								AmountWithFee: nil,
								AmountNoFee:   nil,
								Fee:           nil,
								CountWithFee:  0,
								CountNoFee:    0,
							},
							{
								Id:            2,
								AmountWithFee: nil,
								AmountNoFee:   nil,
								Fee:           nil,
								CountWithFee:  0,
								CountNoFee:    0,
							},
						},
						DailyStatsCount: 2,
						StatsList: []fe.Stats{
							{
								Index: "",
								Date:  "",
								Stats: &fe.DailyStats{
									Id:            0,
									AmountWithFee: nil,
									AmountNoFee:   nil,
									Fee:           types.NewCoins(types.NewCoin("ovg", types.NewInt(10000_0000))),
									CountWithFee:  0,
									CountNoFee:    0,
								},
							},
							{
								Index: "",
								Date:  "",
								Stats: &fe.DailyStats{
									Id:            0,
									AmountWithFee: nil,
									AmountNoFee:   nil,
									Fee:           types.NewCoins(types.NewCoin("ovg", types.NewInt(10000_0000))),
									CountWithFee:  0,
									CountNoFee:    0,
								},
							},
						},
						TariffsList: []fe.Tariffs{
							{
								Denom: "",
								Tariffs: []*fe.Tariff{
									{
										Id:            0,
										Amount:        "",
										Denom:         "",
										MinRefBalance: "",
										Fees: []*fe.Fees{
											{
												AmountFrom:  "",
												Fee:         "",
												RefReward:   "",
												StakeReward: "",
												MinAmount:   0,
												NoRefReward: false,
											},
										},
									},
								},
								Creator: "",
							},
							{
								Denom: "",
								Tariffs: []*fe.Tariff{
									{
										Id:            0,
										Amount:        "",
										Denom:         "",
										MinRefBalance: "",
										Fees: []*fe.Fees{
											{
												AmountFrom:  "",
												Fee:         "",
												RefReward:   "",
												StakeReward: "",
												MinAmount:   0,
												NoRefReward: false,
											},
										},
									},
								},
								Creator: "",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, msg := range tt.args.msg {
				if err := d.Datastore.FeeExluder.InsertToGenesisState(msg); (err != nil) != tt.wantErr {
					t.Errorf("InsertToGenesisState() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
