package wallets

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v3/modules"

	"github.com/forbole/bdjuno/v3/database"
	"github.com/forbole/bdjuno/v3/database/vipcoin/chain/accounts"
	"github.com/forbole/bdjuno/v3/database/vipcoin/chain/wallets"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
	_ modules.GenesisModule            = &Module{}
	_ modules.MessageModule            = &Module{}
)

// Module represents the x/wallets module
type Module struct {
	cdc          codec.Codec
	db           *database.Db
	walletsRepo  wallets.Repository
	accountsRepo accounts.Repository
}

// NewModule returns a new Module instance
func NewModule(cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		cdc:          cdc,
		db:           db,
		walletsRepo:  *wallets.NewRepository(db.Sqlx, cdc),
		accountsRepo: *accounts.NewRepository(db.Sqlx, cdc),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "vipcoin_wallets"
}
