package feeexluder

import (
	"testing"

	"git.ooo.ua/vipcoin/lib/filter"
	fe "git.ooo.ua/vipcoin/ovg-chain/x/feeexcluder/types"
	"github.com/brianvoe/gofakeit/v6"

	d "github.com/forbole/bdjuno/v4/_tests/database"
)

// ERROR: insert or update on table "overgold_feeexcluder_delete_tariffs" violates foreign key constraint "overgold_feeexcluder_delete_tariffs_tariff_id_fkey"

func TestRepository_InsertToMsgDeleteTariffs(t *testing.T) {
	type args struct {
		msg  []fe.MsgDeleteTariffs
		hash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] InsertToMsgDeleteTariffs",
			args: args{
				msg: []fe.MsgDeleteTariffs{
					{
						Creator:  d.TestAddressCreator,
						Denom:    "ovg",
						TariffID: "0",
						FeeID:    "0",
					},
					{
						Creator:  d.TestAddressCreator,
						Denom:    "ovg",
						TariffID: "1",
						FeeID:    "2",
					},
				},
				hash: gofakeit.LetterN(64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := d.Datastore.FeeExluder.InsertToMsgDeleteTariffs(tt.args.hash, tt.args.msg...); (err != nil) != tt.wantErr {
				t.Errorf("InsertToMsgDeleteTariffs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetAllMsgDeleteTariffs(t *testing.T) {
	type args struct {
		filter filter.Filter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] GetAllMsgDeleteTariffs",
			args: args{
				filter: filter.NewFilter(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity, err := d.Datastore.FeeExluder.GetAllMsgDeleteTariffs(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllMsgDeleteTariffs() error = %v, wantErr %v", err, tt.wantErr)
			}

			t.Logf("size: %d", len(entity))
		})
	}
}

func TestRepository_UpdateMsgDeleteTariffs(t *testing.T) {
	type args struct {
		msg  []fe.MsgDeleteTariffs
		id   uint64
		hash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] DeleteMsgDeleteTariffs",
			args: args{
				msg: []fe.MsgDeleteTariffs{
					{
						Creator:  d.TestAddressCreator,
						Denom:    "ovg",
						TariffID: "1",
						FeeID:    "2",
					},
				},
				id:   1,
				hash: gofakeit.LetterN(64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := d.Datastore.FeeExluder.UpdateMsgDeleteTariffs(tt.args.hash, tt.args.id, tt.args.msg...); (err != nil) != tt.wantErr {
				t.Errorf("DeleteMsgDeleteTariffs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_DeleteMsgDeleteTariffs(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] DeleteMsgDeleteTariffs (1)",
			args: args{
				id: 1,
			},
		},
		{
			name: "[success] DeleteMsgDeleteTariffs (2)",
			args: args{
				id: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := d.Datastore.FeeExluder.DeleteMsgDeleteTariffs(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteMsgDeleteTariffs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
