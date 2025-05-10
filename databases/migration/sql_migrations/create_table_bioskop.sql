-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS bioskop
(
    id SERIAL PRIMARY KEY,
    nama VARCHAR(50) NOT NULL,
    lokasi VARCHAR(50) NOT NULL,
    rating DOUBLE PRECISION
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS bioskop;

-- +migrate StatementEnd