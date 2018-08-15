-- +goose Up
CREATE TABLE eth_rates (
    id int NOT NULL AUTO_INCREMENT,
    symbol text NOT NULL,
    price decimal(18, 2) NOT NULL DEFAULT '0',
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE eth_rates;
