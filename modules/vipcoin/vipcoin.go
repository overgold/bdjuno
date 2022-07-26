package vipcoin

import (
	"sync"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v3/logging"
	"github.com/forbole/juno/v3/modules"
	jmodules "github.com/forbole/juno/v3/modules"

	"github.com/forbole/bdjuno/v3/database"
	"github.com/forbole/bdjuno/v3/database/vipcoin/chain/last_block"
	"github.com/forbole/bdjuno/v3/modules/vipcoin/chain/accounts"
	"github.com/forbole/bdjuno/v3/modules/vipcoin/chain/assets"
	"github.com/forbole/bdjuno/v3/modules/vipcoin/chain/banking"
	"github.com/forbole/bdjuno/v3/modules/vipcoin/chain/wallets"
)

var (
	_ modules.Module        = &module{}
	_ modules.GenesisModule = &module{}
)

type vipcoinModule interface {
	jmodules.Module
	jmodules.GenesisModule
	jmodules.MessageModule
}

type module struct {
	cdc            codec.Codec
	db             *database.Db
	lastBlockRepo  last_block.Repository
	logger         logging.Logger
	vipcoinModules []vipcoinModule

	schedulerRun bool
	mutex        sync.RWMutex
}

func NewModule(
	cdc codec.Codec,
	db *database.Db,
	logger logging.Logger,
) *module {
	module := &module{
		cdc:           cdc,
		db:            db,
		lastBlockRepo: *last_block.NewRepository(db.Sqlx),
		logger:        logger,
		vipcoinModules: []vipcoinModule{
			accounts.NewModule(cdc, db),
			assets.NewModule(cdc, db),
			banking.NewModule(cdc, db),
			wallets.NewModule(cdc, db),
		},
	}

	go module.scheduler()

	return module
}

// Name implements modules.Module
func (m *module) Name() string {
	return "vipcoin"
}
