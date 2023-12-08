package feeexcluder

import (
	"encoding/json"

	feeexluder "git.ooo.ua/vipcoin/ovg-chain/x/feeexcluder/types"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/rs/zerolog/log"
)

// HandleGenesis implements GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", feeexluder.ModuleName).Msg("parsing genesis")

	// Unmarshal the bank state
	var genesisState feeexluder.GenesisState
	if err := m.cdc.UnmarshalJSON(appState[feeexluder.ModuleName], &genesisState); err != nil {
		return err
	}

	return m.feeexcluderRepo.InsertToGenesisState(genesisState)
}
