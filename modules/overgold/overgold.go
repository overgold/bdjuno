package overgold

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v5/logging"
	jmodules "github.com/forbole/juno/v5/modules"
	"github.com/forbole/juno/v5/node"

	"github.com/forbole/bdjuno/v4/database/overgold/chain/last_block"

	"github.com/forbole/bdjuno/v4/database"
	"github.com/forbole/bdjuno/v4/modules/overgold/chain/allowed"
	overgoldAllowedSource "github.com/forbole/bdjuno/v4/modules/overgold/chain/allowed/source"
	customBank "github.com/forbole/bdjuno/v4/modules/overgold/chain/bank"
	overgoldBankSource "github.com/forbole/bdjuno/v4/modules/overgold/chain/bank/source"
	"github.com/forbole/bdjuno/v4/modules/overgold/chain/core"
	overgoldCoreSource "github.com/forbole/bdjuno/v4/modules/overgold/chain/core/source"
	"github.com/forbole/bdjuno/v4/modules/overgold/chain/feeexcluder"
	overgoldFeeExcluderSource "github.com/forbole/bdjuno/v4/modules/overgold/chain/feeexcluder/source"
	"github.com/forbole/bdjuno/v4/modules/overgold/chain/referral"
	overgoldReferralSource "github.com/forbole/bdjuno/v4/modules/overgold/chain/referral/source"
	"github.com/forbole/bdjuno/v4/modules/overgold/chain/stake"
	overgoldStakeSource "github.com/forbole/bdjuno/v4/modules/overgold/chain/stake/source"
)

var (
	_ jmodules.Module        = &Module{}
	_ jmodules.GenesisModule = &Module{}
)

type overgoldModule interface {
	jmodules.Module
	jmodules.GenesisModule
	jmodules.MessageModule
}

type Module struct {
	cdc             codec.Codec
	db              *database.Db
	lastBlockRepo   last_block.Repository
	logger          logging.Logger
	overgoldModules []overgoldModule
	node            node.Node
}

func NewModule(
	cdc codec.Codec,
	db *database.Db,
	node node.Node,
	logger logging.Logger,

	overGoldAllowedSource overgoldAllowedSource.Source,
	overGoldBankSource overgoldBankSource.Source,
	overGoldCoreSource overgoldCoreSource.Source,
	overGoldFeeExcluderSource overgoldFeeExcluderSource.Source,
	overGoldReferralSource overgoldReferralSource.Source,
	overGoldStakeSource overgoldStakeSource.Source,
) *Module {
	module := &Module{
		cdc:           cdc,
		db:            db,
		lastBlockRepo: *last_block.NewRepository(db.Sqlx),
		node:          node,
		logger:        logger,
		overgoldModules: []overgoldModule{
			// OverGold modules
			allowed.NewModule(overGoldAllowedSource, cdc, db),
			core.NewModule(overGoldCoreSource, cdc, db),
			feeexcluder.NewModule(overGoldFeeExcluderSource, cdc, db),
			referral.NewModule(overGoldReferralSource, cdc, db),
			stake.NewModule(overGoldStakeSource, cdc, db),

			// custom SDK modules
			customBank.NewModule(overGoldBankSource, cdc, db),
		},
	}

	go module.scheduler()

	return module
}

// Name implements modules.Module
func (m *Module) Name() string {
	return module
}
