package assets

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v3/modules"

	"github.com/forbole/bdjuno/v3/database"
	"github.com/forbole/bdjuno/v3/database/vipcoin/chain/assets"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represents the x/assets module
type Module struct {
	assetRepo assets.Repository
	cdc       codec.Codec
	db        *database.Db
}

// NewModule returns a new Module instance
func NewModule(cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		assetRepo: *assets.NewRepository(db.Sqlx, cdc),
		cdc:       cdc,
		db:        db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "vipcoin_assets"
}
