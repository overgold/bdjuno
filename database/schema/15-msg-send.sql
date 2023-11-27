-- +migrate Up
CREATE TABLE IF NOT EXISTS msg_send
(
    id               BIGSERIAL      NOT NULL PRIMARY KEY,
    tx_hash          TEXT           NOT NULL,
    from_address     TEXT           NOT NULL,
    to_address       TEXT           NOT NULL,
    amount           BIGINT[]       NOT NULL,
    denom            TEXT[]         NOT NULL
);

CREATE TABLE IF NOT EXISTS msg_multi_send_data
(
    id               BIGSERIAL      NOT NULL PRIMARY KEY,
    address          TEXT           NOT NULL,
    amount           BIGINT[]       NOT NULL,
    denom            TEXT[]         NOT NULL
);

CREATE TABLE IF NOT EXISTS msg_multi_send
(
    id               BIGSERIAL      NOT NULL PRIMARY KEY,
    tx_hash          TEXT           NOT NULL,
    input_ids        BIGSERIAL[]    NOT NULL, -- references to 'msg_multi_send_data'
    output_ids       BIGSERIAL[]    NOT NULL  -- references to 'msg_multi_send_data'
);


-- +migrate Down
DROP TABLE IF EXISTS msg_multi_send CASCADE;
DROP TABLE IF EXISTS msg_multi_send_data CASCADE;
DROP TABLE IF EXISTS overgold_allowed_update_addresses CASCADE;