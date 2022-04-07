-- +migrate Up
CREATE TABLE IF NOT EXISTS vipcoin_chain_banking_base_transfers (
    id                  SERIAL PRIMARY KEY NOT NULL,            -- banking id
    asset               TEXT NOT NULL,                          -- banking asset
    amount              BIGSERIAL,                              -- banking amount
    kind                INT,                                    -- banking kind
    extras              JSONB,                                  -- banking extras
    timestamp           TIMESTAMP WITHOUT TIME ZONE NOT NULL,   -- banking timestamp
    tx_hash             TEXT NOT NULL                           -- banking tx hash
);

CREATE TABLE IF NOT EXISTS vipcoin_chain_banking_system_transfer (
    id                  SERIAL UNIQUE PRIMARY KEY NOT NULL,     -- banking id
    wallet_from         TEXT NOT NULL,                          -- banking wallet from
    wallet_to           TEXT NOT NULL,                          -- banking wallet to
    CONSTRAINT system_transfer_transfers_id_pkey FOREIGN KEY(id)
        REFERENCES vipcoin_chain_banking_base_transfers(id)
);

CREATE TABLE IF NOT EXISTS vipcoin_chain_banking_payment (
    id                  SERIAL UNIQUE PRIMARY KEY NOT NULL,     -- banking id
    wallet_from         TEXT NOT NULL,                          -- banking wallet from
    wallet_to           TEXT NOT NULL,                          -- banking wallet to
    fee                 BIGSERIAL,                              -- banking fee
    CONSTRAINT payment_transfers_id_pkey FOREIGN KEY(id)
        REFERENCES vipcoin_chain_banking_base_transfers(id)
);

CREATE TABLE IF NOT EXISTS vipcoin_chain_banking_withdraw (
    id                  SERIAL UNIQUE PRIMARY KEY NOT NULL,     -- banking id
    wallet              TEXT NOT NULL,                          -- banking wallet
    CONSTRAINT withdraw_transfers_id_pkey FOREIGN KEY(id)
        REFERENCES vipcoin_chain_banking_base_transfers(id)
);

CREATE TABLE IF NOT EXISTS vipcoin_chain_banking_issue (
    id                  SERIAL UNIQUE PRIMARY KEY NOT NULL,     -- banking id
    wallet              TEXT NOT NULL,                          -- banking wallet
    CONSTRAINT issue_transfers_id_pkey FOREIGN KEY(id)
        REFERENCES vipcoin_chain_banking_base_transfers(id)
);

CREATE TABLE IF NOT EXISTS vipcoin_chain_banking_msg_system_transfer (
    creator             TEXT       NOT NULL,                    -- banking creator
    wallet_from         TEXT       NOT NULL,                    -- banking wallet from
    wallet_to           TEXT       NOT NULL,                    -- banking wallet to
    asset               TEXT       NOT NULL,                    -- banking asset
    amount              BIGINT,                                 -- banking amount
    extras              JSONB                                   -- banking extras
);

CREATE TABLE IF NOT EXISTS vipcoin_chain_banking_system_reward_transfer (
    creator             TEXT NOT NULL,                          -- banking creator
    wallet_from         TEXT NOT NULL,                          -- banking wallet from
    wallet_to           TEXT NOT NULL,                          -- banking wallet to
    asset               TEXT NOT NULL,                          -- banking asset
    amount              BIGINT,                                 -- banking amount
    extras              JSONB                                   -- banking extras
);

CREATE TABLE IF NOT EXISTS vipcoin_chain_banking_msg_payment (
    creator             TEXT NOT NULL,                          -- banking creator
    wallet_from         TEXT NOT NULL,                          -- banking wallet from
    wallet_to           TEXT NOT NULL,                          -- banking wallet to
    asset               TEXT NOT NULL,                          -- banking asset
    amount              BIGINT,                                 -- banking amount
    extras              JSONB                                   -- banking extras
);

CREATE TABLE IF NOT EXISTS vipcoin_chain_banking_msg_withdraw (
    creator             TEXT NOT NULL,                          -- banking creator
    wallet              TEXT NOT NULL,                          -- banking wallet
    asset               TEXT NOT NULL,                          -- banking asset
    amount              BIGINT,                                 -- banking amount
    extras              JSONB                                   -- banking extras
);

CREATE TABLE IF NOT EXISTS vipcoin_chain_banking_msg_issue (
    creator             TEXT NOT NULL,                          -- banking creator
    wallet              TEXT NOT NULL,                          -- banking wallet
    asset               TEXT NOT NULL,                          -- banking asset
    amount              BIGINT,                                 -- banking amount
    extras              JSONB                                   -- banking extras
);

CREATE TABLE IF NOT EXISTS vipcoin_chain_banking_set_transfer_extra (
    creator             TEXT NOT NULL,                          -- banking creator
    id                  BIGINT,                                 -- banking id
    extras              JSONB                                   -- banking extras
);

CREATE TABLE IF NOT EXISTS vipcoin_chain_banking_set_reward_manager_address (
    creator             TEXT NOT NULL,                          -- banking creator
    address             TEXT NOT NULL                           -- banking address
);

-- +migrate Down
DROP TABLE IF EXISTS vipcoin_chain_banking_system_transfer CASCADE;
DROP TABLE IF EXISTS vipcoin_chain_banking_system_reward_transfer CASCADE;
DROP TABLE IF EXISTS vipcoin_chain_banking_payment CASCADE;
DROP TABLE IF EXISTS vipcoin_chain_banking_withdraw CASCADE;
DROP TABLE IF EXISTS vipcoin_chain_banking_issue CASCADE;
DROP TABLE IF EXISTS vipcoin_chain_banking_set_transfer_extra CASCADE;
DROP TABLE IF EXISTS vipcoin_chain_banking_set_reward_manager_address CASCADE;
DROP TABLE IF EXISTS vipcoin_chain_banking_msg_system_transfer CASCADE;
DROP TABLE IF EXISTS vipcoin_chain_banking_msg_payment CASCADE;
DROP TABLE IF EXISTS vipcoin_chain_banking_msg_withdraw CASCADE;
DROP TABLE IF EXISTS vipcoin_chain_banking_msg_issue CASCADE;
