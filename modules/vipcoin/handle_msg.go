package vipcoin

import (
	"errors"
	"fmt"

	"git.ooo.ua/vipcoin/lib/errs"
	"git.ooo.ua/vipcoin/lib/filter"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v2/modules"
	"github.com/forbole/juno/v2/types"

	dbtypes "github.com/forbole/bdjuno/v2/database/types"
)

// HandleMsg implements MessageModule
func (m *module) HandleMsg(index int, msg sdk.Msg, tx *types.Tx) error {
	m.mutex.RLock()
	if m.schedulerRun {
		m.mutex.RUnlock()
		return nil
	}
	m.mutex.RUnlock()

	m.mutex.Lock()
	m.schedulerRun = true
	go m.parseBlock()
	m.mutex.Unlock()

	return nil
}

func (m *module) parseBlock() {
	defer func() {
		m.mutex.Lock()
		m.schedulerRun = false
		m.mutex.Unlock()
	}()

	last_block, err := m.lastBlockRepo.Get()
	if err != nil {
		m.logger.Error("Fail lastBlockRepo.Get", "module", "vipcoin", "error", err)
		return
	}

	for {
		last_block += 1

		block, err := m.db.GetBlock(filter.NewFilter().SetArgument(dbtypes.FieldHeight, last_block))
		if err != nil {
			if errors.As(err, &errs.NotFound{}) {
				return
			}

			m.logger.Error("Fail GetBlock", "module", "vipcoin", "error", err)
			return
		}

		m.logger.Debug("parse block", "height", block.Height)

		txs, err := m.db.GetTransactions(
			filter.NewFilter().
				SetCondition(filter.ConditionAND).
				SetArgument(dbtypes.FieldHeight, block.Height).
				SetArgument(dbtypes.FieldSuccess, true))
		if err != nil {
			if !errors.As(err, &errs.NotFound{}) {
				m.logger.Error("Fail GetTransactions", "module", "vipcoin", "error", err)
			}

			if err := m.lastBlockRepo.Update(last_block); err != nil {
				m.logger.Error("Fail lastBlockRepo.Update", "module", "vipcoin", "error", err)
				return
			}
			continue
		}

		for _, tx := range txs {
			if err := m.parseMessages(tx); err != nil {
				m.logger.Error("Fail parseMessages", "module", "vipcoin", "error", err)
			}
		}

		if err := m.lastBlockRepo.Update(last_block); err != nil {
			m.logger.Error("Fail lastBlockRepo.Update", "module", "vipcoin", "error", err)
			return
		}
	}
}

// parseMessages - parse messages from transaction
func (m *module) parseMessages(tx *types.Tx) error {
	for i, msg := range tx.Body.Messages {
		var stdMsg sdk.Msg
		if err := m.cdc.UnpackAny(msg, &stdMsg); err != nil {
			return fmt.Errorf("error while unpacking message: %s", err)
		}

		for _, module := range m.vipcoinModules {
			if messageModule, ok := module.(modules.MessageModule); ok {
				if err := messageModule.HandleMsg(i, stdMsg, tx); err != nil {
					m.logger.MsgError(module, tx, stdMsg, err)
				}
			}
		}
	}

	return nil
}
