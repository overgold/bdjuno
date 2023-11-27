-- +migrate Up
CREATE TYPE SEND_DATA AS
(
    address          TEXT,
    coins            COIN[]
);

CREATE TABLE IF NOT EXISTS msg_multi_send
(
    id               BIGSERIAL      NOT NULL PRIMARY KEY,
    tx_hash          TEXT           NOT NULL,
<<<<<<< HEAD
    inputs           SEND_DATA[]    NOT NULL DEFAULT '{}',
    outputs          SEND_DATA[]    NOT NULL DEFAULT '{}'
=======
    inputs           COIN[]         NOT NULL,
    outputs          COIN[]         NOT NULL
>>>>>>> 4ff886e (VC-14559: added bank module (send/multi_send))
);

CREATE TABLE IF NOT EXISTS msg_send
(
    id               BIGSERIAL      NOT NULL PRIMARY KEY,
    tx_hash          TEXT           NOT NULL,
    from_address     TEXT           NOT NULL,
    to_address       TEXT           NOT NULL,
    amount           COIN[]         NOT NULL DEFAULT '{}'
);

-- +migrate Down
DROP TABLE IF EXISTS msg_send CASCADE;
DROP TABLE IF EXISTS msg_multi_send CASCADE;
DROP TYPE IF EXISTS SEND_DATA CASCADE;