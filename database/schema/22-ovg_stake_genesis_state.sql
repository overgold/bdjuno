-- +migrate Up
CREATE TABLE IF NOT EXISTS overgold_stake_transaction
(
    id               BIGSERIAL  NOT NULL PRIMARY KEY,
    kind             INTEGER    NOT NULL,
    address_from     TEXT       NOT NULL,
    address_to       TEXT       NOT NULL,
    description      TEXT       NOT NULL,
    timestamp        TIMESTAMP  NOT NULL,
    amount           COIN[]     NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_sell_stake_params
(
    id                  BIGSERIAL  NOT NULL PRIMARY KEY,
    min_sell_requests   BIGINT     NOT NULL,
    max_sell_requests   BIGINT     NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_unique_users
(
    id          BIGSERIAL   NOT NULL PRIMARY KEY,
    burn        TEXT[]      NOT NULL,
    buy         TEXT[]      NOT NULL,
    issue       TEXT[]      NOT NULL,
    sell        TEXT[]      NOT NULL,
    total       TEXT[]      NOT NULL,
    withdraw    TEXT[]      NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_min_amounts
(
    id                  BIGSERIAL  NOT NULL PRIMARY KEY,
    min_sell_requests   COIN[]     NOT NULL,
    max_sell_requests   COIN[]     NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_stake
(
    id              BIGSERIAL  NOT NULL PRIMARY KEY,
    reward_amount   COIN[]     NOT NULL,
    sell_amount     COIN[]     NOT NULL,
    stake_amount    COIN[]     NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_daily_stats
(
    id                      BIGSERIAL   NOT NULL PRIMARY KEY,
    msg_id                  BIGSERIAL   NOT NULL,
    count_burn              BIGINT      NOT NULL,
    count_buy               BIGINT      NOT NULL,
    count_issue             BIGINT      NOT NULL,
    count_sell              BIGINT      NOT NULL,
    count_withdraw          BIGINT      NOT NULL,
    count_users             BIGINT      NOT NULL,
    count_users_burn        BIGINT      NOT NULL,
    count_users_buy         BIGINT      NOT NULL,
    count_users_issue       BIGINT      NOT NULL,
    count_users_sell        BIGINT      NOT NULL,
    count_users_withdraw    BIGINT      NOT NULL,
    distributions_by_user   BIGINT      NOT NULL,
    amount_burn             COIN[]      NOT NULL,
    amount_buy              COIN[]      NOT NULL,
    amount_issue            COIN[]      NOT NULL,
    amount_sell             COIN[]      NOT NULL,
    amount_withdraw         COIN[]      NOT NULL,
    distributions_total     COIN[]      NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_sell_params
(
    id                  BIGSERIAL  NOT NULL PRIMARY KEY,
    stake_params_id     BIGSERIAL            REFERENCES overgold_stake_sell_stake_params(id),
    creator             TEXT       NOT NULL,
    denom               TEXT       NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_user_stats
(
    id                  BIGSERIAL   NOT NULL PRIMARY KEY,
    unique_users_id     BIGINT      NOT NULL REFERENCES overgold_stake_unique_users(id),
    date                TIMESTAMP   NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_stake_limits
(
    id              BIGSERIAL  NOT NULL PRIMARY KEY,
    min_amounts_id  BIGSERIAL            REFERENCES overgold_stake_min_amounts(id),
    creator         TEXT       NOT NULL,
    denom           TEXT       NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_sell_requests
(
    id               BIGSERIAL  NOT NULL PRIMARY KEY,
    account_address  TEXT       NOT NULL,
    amount           COIN[]     NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_buy_requests
(
    id               BIGSERIAL  NOT NULL PRIMARY KEY,
    account_address  TEXT       NOT NULL,
    amount           COIN[]     NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_stakes
(
    id               BIGSERIAL  NOT NULL PRIMARY KEY,
    stake_id         BIGSERIAL            REFERENCES overgold_stake_stake(id),
    account_address  TEXT       NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_stats
(
    id          BIGSERIAL   NOT NULL PRIMARY KEY,
    stats_id    BIGSERIAL            REFERENCES overgold_stake_daily_stats(id),
    date        TIMESTAMP   NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_genesis_state
(
    id                  BIGSERIAL   NOT NULL PRIMARY KEY,
    daily_stats_count   BIGINT      NOT NULL,
    total_free          BIGINT      NOT NULL,
    total_sell          BIGINT      NOT NULL,
    users_count         BIGINT      NOT NULL
);

CREATE TABLE IF NOT EXISTS overgold_stake_m2m_genesis_state_stats
(
    genesis_state_id BIGSERIAL REFERENCES overgold_stake_genesis_state(id),
    stats_id         BIGSERIAL REFERENCES overgold_stake_stats(id),
    PRIMARY KEY (genesis_state_id, stats_id)
);

CREATE TABLE IF NOT EXISTS overgold_stake_m2m_genesis_state_stakes
(
    genesis_state_id BIGSERIAL REFERENCES overgold_stake_genesis_state(id),
    stakes_id          BIGSERIAL REFERENCES overgold_stake_stakes(id),
    PRIMARY KEY (genesis_state_id, stakes_id)
);

CREATE TABLE IF NOT EXISTS overgold_stake_m2m_genesis_state_buy_requests
(
    genesis_state_id BIGSERIAL REFERENCES overgold_stake_genesis_state(id),
    buy_requests_id  BIGSERIAL REFERENCES overgold_stake_buy_requests(id),
    PRIMARY KEY (genesis_state_id, buy_requests_id)
);

CREATE TABLE IF NOT EXISTS overgold_stake_m2m_genesis_state_sell_requests
(
    genesis_state_id BIGSERIAL REFERENCES overgold_stake_genesis_state(id),
    sell_requests_id BIGSERIAL REFERENCES overgold_stake_sell_requests(id),
    PRIMARY KEY (genesis_state_id, sell_requests_id)
);

CREATE TABLE IF NOT EXISTS overgold_stake_m2m_genesis_state_stake_limits
(
    genesis_state_id BIGSERIAL REFERENCES overgold_stake_genesis_state(id),
    state_limits_id  BIGSERIAL REFERENCES overgold_stake_stake_limits(id),
    PRIMARY KEY (genesis_state_id, state_limits_id)
);

CREATE TABLE IF NOT EXISTS overgold_stake_m2m_genesis_state_user_stats
(
    genesis_state_id BIGSERIAL REFERENCES overgold_stake_genesis_state(id),
    user_stats_id    BIGSERIAL REFERENCES overgold_stake_user_stats(id),
    PRIMARY KEY (genesis_state_id, user_stats_id)
);

CREATE TABLE IF NOT EXISTS overgold_stake_m2m_genesis_state_sell_params
(
    genesis_state_id BIGSERIAL REFERENCES overgold_stake_genesis_state(id),
    sell_params_id   BIGSERIAL REFERENCES overgold_stake_sell_params(id),
    PRIMARY KEY (genesis_state_id, sell_params_id)
);

CREATE TABLE IF NOT EXISTS overgold_stake_m2m_stake_transaction
(
    stake_id        BIGSERIAL REFERENCES overgold_stake_stake(id),
    transaction_id  BIGSERIAL REFERENCES overgold_stake_transaction(id),
    PRIMARY KEY (stake_id, transaction_id)
);

-- +migrate Down

DROP TABLE IF EXISTS overgold_stake_m2m_stake_transaction CASCADE;
DROP TABLE IF EXISTS overgold_stake_m2m_genesis_state_sell_params CASCADE;
DROP TABLE IF EXISTS overgold_stake_m2m_genesis_state_user_stats CASCADE;
DROP TABLE IF EXISTS overgold_stake_m2m_genesis_state_stake_limits CASCADE;
DROP TABLE IF EXISTS overgold_stake_m2m_genesis_state_sell_requests CASCADE;
DROP TABLE IF EXISTS overgold_stake_m2m_genesis_state_buy_requests CASCADE;
DROP TABLE IF EXISTS overgold_stake_m2m_genesis_state_stakes CASCADE;
DROP TABLE IF EXISTS overgold_stake_m2m_genesis_state_stats CASCADE;
DROP TABLE IF EXISTS overgold_stake_m2m_genesis_state_stats CASCADE;
DROP TABLE IF EXISTS overgold_stake_genesis_state CASCADE;
DROP TABLE IF EXISTS overgold_stake_stats CASCADE;
DROP TABLE IF EXISTS overgold_stake_stakes CASCADE;
DROP TABLE IF EXISTS overgold_stake_buy_requests CASCADE;
DROP TABLE IF EXISTS overgold_stake_sell_requests CASCADE;
DROP TABLE IF EXISTS overgold_stake_stake_limits CASCADE;
DROP TABLE IF EXISTS overgold_stake_user_stats CASCADE;
DROP TABLE IF EXISTS overgold_stake_sell_params CASCADE;
DROP TABLE IF EXISTS overgold_stake_daily_stats CASCADE;
DROP TABLE IF EXISTS overgold_stake_stake CASCADE;
DROP TABLE IF EXISTS overgold_stake_min_amounts CASCADE;
DROP TABLE IF EXISTS overgold_stake_unique_users CASCADE;
DROP TABLE IF EXISTS overgold_stake_sell_stake_params CASCADE;
DROP TABLE IF EXISTS overgold_stake_transaction CASCADE;
