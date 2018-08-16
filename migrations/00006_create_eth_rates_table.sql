-- +goose Up
CREATE TABLE eth_rates (
    id int NOT NULL AUTO_INCREMENT,
    symbol varchar(10) NOT NULL,
    price decimal(18, 2) NOT NULL DEFAULT '0',
    time int(11) NOT NULL,
    PRIMARY KEY(id),
    KEY (symbol, time)
);

-- +goose Down
DROP TABLE eth_rates;
