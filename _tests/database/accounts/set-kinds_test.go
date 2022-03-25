/*
 * Copyright 2022 Business Process Technologies. All rights reserved.
 */

package accounts

import (
	"reflect"
	"testing"

	accountstypes "git.ooo.ua/vipcoin/chain/x/accounts/types"
	"git.ooo.ua/vipcoin/lib/filter"
	"github.com/cosmos/cosmos-sdk/simapp"
	accountsdb "github.com/forbole/bdjuno/v2/database/vipcoin/chain/accounts"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestRepository_SaveSetKinds(t *testing.T) {
	// db, err := sqlx.Connect("pgx", "host=10.10.1.79 port=5432 user=postgres dbname=juno password=postgres sslmode=disable")
	db, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=juno password=postgres sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// Create the codec
	codec := simapp.MakeTestEncodingConfig()

	type args struct {
		msg []*accountstypes.MsgSetKinds
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				msg: []*accountstypes.MsgSetKinds{
					{
						Creator: "vcg1ljs7p2p9ae3en8knr3d3ke8srsfcj2zjvefv0g",
						Hash:    "a935ea2c467d7f666ea2a67870564f2efb902c05f0a2bb4b6202832aedd26cd1",
						Kinds:   []accountstypes.AccountKind{accountstypes.ACCOUNT_KIND_SYSTEM},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := accountsdb.NewRepository(db, codec.Marshaler)

			if err := r.SaveSetKinds(tt.args.msg...); (err != nil) != tt.wantErr {
				t.Errorf("Repository.SaveSetKinds() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetSetKinds(t *testing.T) {
	// db, err := sqlx.Connect("pgx", "host=10.10.1.79 port=5432 user=postgres dbname=juno password=postgres sslmode=disable")
	db, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=juno password=postgres sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// Create the codec
	codec := simapp.MakeTestEncodingConfig()

	type args struct {
		accfilter filter.Filter
	}
	tests := []struct {
		name    string
		args    args
		want    []*accountstypes.MsgSetKinds
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				accfilter: filter.NewFilter().SetArgument("creator", "vcg1ljs7p2p9ae3en8knr3d3ke8srsfcj2zjvefv0g"),
			},
			want: []*accountstypes.MsgSetKinds{
				{
					Creator: "vcg1ljs7p2p9ae3en8knr3d3ke8srsfcj2zjvefv0g",
					Hash:    "a935ea2c467d7f666ea2a67870564f2efb902c05f0a2bb4b6202832aedd26cd1",
					Kinds:   []accountstypes.AccountKind{accountstypes.ACCOUNT_KIND_SYSTEM},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := accountsdb.NewRepository(db, codec.Marshaler)

			got, err := r.GetSetKinds(tt.args.accfilter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetSetKinds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetSetKinds() = %v, want %v", got, tt.want)
			}
		})
	}
}
