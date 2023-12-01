package referral

import (
	"testing"

	"git.ooo.ua/vipcoin/lib/filter"
	referral "git.ooo.ua/vipcoin/ovg-chain/x/referral/types"
	"github.com/brianvoe/gofakeit/v6"

	d "github.com/forbole/bdjuno/v4/_tests/database"
	"github.com/forbole/bdjuno/v4/database/types"
)

func TestRepository_InsertMsgSetRefferrer(t *testing.T) {
	type args struct {
		msg  []referral.MsgSetReferrer
		hash string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] InsertMsgSetReferrer",
			args: args{
				msg: []referral.MsgSetReferrer{
					{
						Creator:         d.TestAddressCreator,
						ReferrerAddress: d.TestAddress,
						ReferralAddress: d.TestAddressCreator,
					},
				},
				hash: gofakeit.LetterN(64),
			},
		},
		{
			name: "[success] InsertMsgSetReferrer (random)",
			args: args{
				msg: []referral.MsgSetReferrer{
					{
						Creator:         d.TestAddressCreator,
						ReferrerAddress: gofakeit.Regex("^ovg[a-z0-9]{39}"),
						ReferralAddress: gofakeit.Regex("^ovg[a-z0-9]{39}"),
					},
				},
				hash: gofakeit.LetterN(64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := d.Datastore.Referral.InsertMsgSetReferrer(tt.args.hash, tt.args.msg...)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertMsgSetReferrer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetAllMsgSetRefferrer(t *testing.T) {
	type args struct {
		filter filter.Filter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "[success] GetAllMsgSetRefferrer",
			args: args{
				filter: filter.NewFilter(),
			},
		},
		{
			name: "[success] GetAllMsgSetRefferrer by referrer",
			args: args{
				filter: filter.NewFilter().SetArgument(types.FieldReferrerAddress, d.TestAddress),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entities, err := d.Datastore.Referral.GetAllMsgSetReferrer(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllMsgSetRefferrer() error = %v, wantErr %v", err, tt.wantErr)
			}

			t.Logf("size: %d", len(entities))
			for i, e := range entities {
				t.Logf("%d: referrer: %s", i, e.ReferrerAddress)
			}
		})
	}
}
