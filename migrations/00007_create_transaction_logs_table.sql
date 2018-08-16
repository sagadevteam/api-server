-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE transaction_logs (
    transaction_log_id int NOT NULL AUTO_INCREMENT,
    inventory_id int,
    buyer_id int,
    buy_time int(11),
    PRIMARY KEY(transaction_log_id),
    FOREIGN KEY (buyer_id) REFERENCES users(user_id),
    FOREIGN KEY (inventory_id) REFERENCES inventories(inventory_id),
    KEY (buy_time)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
-- +goose Down
DROP TABLE transaction_logs;