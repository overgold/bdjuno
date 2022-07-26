package accounts

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v3/modules"

	"github.com/forbole/bdjuno/v3/database"
	"github.com/forbole/bdjuno/v3/database/vipcoin/chain/accounts"
	"github.com/forbole/bdjuno/v3/database/vipcoin/chain/assets"
	"github.com/forbole/bdjuno/v3/database/vipcoin/chain/wallets"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represents the x/accounts module
type Module struct {
	cdc         codec.Codec
	db          *database.Db
	walletsRepo wallets.Repository
	accountRepo accounts.Repository
	assetRepo   assets.Repository
}

// NewModule returns a new Module instance
func NewModule(cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		cdc:         cdc,
		db:          db,
		accountRepo: *accounts.NewRepository(db.Sqlx, cdc),
		walletsRepo: *wallets.NewRepository(db.Sqlx, cdc),
		assetRepo:   *assets.NewRepository(db.Sqlx, cdc),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "vipcoin_accounts"
}
