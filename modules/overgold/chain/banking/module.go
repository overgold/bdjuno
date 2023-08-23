package banking

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v3/modules"

	"github.com/forbole/bdjuno/v3/database"
	"github.com/forbole/bdjuno/v3/database/overgold/chain/accounts"
	"github.com/forbole/bdjuno/v3/database/overgold/chain/assets"
	"github.com/forbole/bdjuno/v3/database/overgold/chain/banking"
	"github.com/forbole/bdjuno/v3/database/overgold/chain/wallets"
	"github.com/forbole/bdjuno/v3/modules/overgold/chain/banking/source"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represents the x/banking module
type Module struct {
	cdc          codec.Codec
	db           *database.Db
	bankingRepo  banking.Repository
	walletsRepo  wallets.Repository
	assetRepo    assets.Repository
	accountsRepo accounts.Repository
	keeper       source.Source
}

// NewModule returns a new Module instance
func NewModule(keeper source.Source, cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		cdc:          cdc,
		db:           db,
		bankingRepo:  *banking.NewRepository(db.Sqlx, cdc),
		walletsRepo:  *wallets.NewRepository(db.Sqlx, cdc),
		assetRepo:    *assets.NewRepository(db.Sqlx, cdc),
		accountsRepo: *accounts.NewRepository(db.Sqlx, cdc),
		keeper:       keeper,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "overgold_banking"
}
