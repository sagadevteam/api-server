-- +goose Up
CREATE TABLE deposits (
    deposit_id int NOT NULL AUTO_INCREMENT,
    txhash varchar(80) NOT NULL,
    address varchar(60) NOT NULL,
    approved tinyint(1) NOT NULL DEFAULT '0',
    PRIMARY KEY(deposit_id),
    KEY (txhash, address)
);

-- +goose Down
DROP TABLE deposits;
