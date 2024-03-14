-- +migrate Up

CREATE TABLE IF NOT EXISTS overgold_stake_create_system_stake_account_address
(
    id      BIGSERIAL NOT NULL PRIMARY KEY,
    tx_hash TEXT      NOT NULL,
    creator TEXT      NOT NULL,
    address TEXT      NOT NULL
);

CREATE UNIQUE INDEX idx_overgold_stake_create_system_stake_account_address ON overgold_stake_create_system_stake_account_address (tx_hash);

CREATE TABLE IF NOT EXISTS overgold_stake_update_system_stake_account_address
(
    id      BIGSERIAL NOT NULL PRIMARY KEY,
    tx_hash TEXT      NOT NULL,
    creator TEXT      NOT NULL,
    address TEXT      NOT NULL
);

CREATE UNIQUE INDEX idx_overgold_stake_update_system_stake_account_address ON overgold_stake_update_system_stake_account_address (tx_hash);


CREATE TABLE IF NOT EXISTS overgold_stake_delete_system_stake_account_address
(
    id      BIGSERIAL NOT NULL PRIMARY KEY,
    tx_hash TEXT      NOT NULL,
    creator TEXT      NOT NULL
);

CREATE UNIQUE INDEX idx_overgold_stake_delete_system_stake_account_address ON overgold_stake_delete_system_stake_account_address (tx_hash);


-- +migrate Down

DROP TABLE IF EXISTS overgold_stake_create_system_stake_account_address;
DROP TABLE IF EXISTS overgold_stake_update_system_stake_account_address;
DROP TABLE IF EXISTS overgold_stake_delete_system_stake_account_address;

DROP INDEX IF EXISTS idx_overgold_stake_create_system_stake_account_address;
DROP INDEX IF EXISTS idx_overgold_stake_update_system_stake_account_address;
DROP INDEX IF EXISTS idx_overgold_stake_delete_system_stake_account_address;




