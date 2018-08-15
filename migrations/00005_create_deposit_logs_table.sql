-- +goose Up
CREATE TABLE deposit_logs (
    log_id int NOT NULL AUTO_INCREMENT,
    txhash text NOT NULL,
    address text NOT NULL,
    approved tinyint(1) NOT NULL DEFAULT '0',
    PRIMARY KEY(log_id)
);

-- +goose Down
DROP TABLE deposit_logs;
