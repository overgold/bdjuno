package core

import (
	"testing"

	"git.ooo.ua/vipcoin/lib/filter"
	stake "git.ooo.ua/vipcoin/ovg-chain/x/stake/types"
	"github.com/brianvoe/gofakeit/v6"

	d "github.com/forbole/bdjuno/v4/_tests/database"
	"github.com/forbole/bdjuno/v4/database/types"
)

func TestRepository_InsertMsgDistributeRewards(t *testing.T) {
	type args struct {
		msg  []stake.MsgDistributeRewards
		hash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] InsertMsgDistributeRewards",
			args: args{
				msg: []stake.MsgDistributeRewards{
					{
						Creator: d.TestAddressCreator,
					},
				},
				hash: gofakeit.LetterN(64),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := d.Datastore.Stake.InsertMsgDistributeRewards(tt.args.hash, tt.args.msg...)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertMsgDistributeRewards() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetAllMsgDistributeRewards(t *testing.T) {
	type args struct {
		filter filter.Filter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] GetAllMsgDistributeRewards",
			args: args{
				filter: filter.NewFilter(),
			},
		},
		{
			name: "[success] GetAllMsgDistributeRewards by address",
			args: args{
				filter: filter.NewFilter().SetArgument(types.FieldAddress, d.TestAddressCreator),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity, err := d.Datastore.Stake.GetAllMsgDistributeRewards(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllMsgDistributeRewards() error = %v, wantErr %v", err, tt.wantErr)
			}

			t.Logf("size: %d", len(entity))
			for _, e := range entity {
				t.Logf("creator: %s", e.Creator)
			}
		})
	}
}
