-- +migrate Up

CREATE TABLE IF NOT EXISTS overgold_stake_manage_system_stake
(
    id      BIGSERIAL NOT NULL PRIMARY KEY,
    tx_hash TEXT      NOT NULL,
    creator TEXT      NOT NULL,
    amount  BIGINT    NOT NULL,
    kind    TEXT      NOT NULL
);

CREATE INDEX idx_overgold_stake_manage_system_stake ON overgold_stake_manage_system_stake (tx_hash);
-- +migrate Down

DROP TABLE IF EXISTS overgold_stake_manage_system_stake;
DROP INDEX IF EXISTS idx_overgold_stake_manage_system_stake;