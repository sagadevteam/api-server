-- +goose Up
CREATE TABLE inventories (
    inventory_id int NOT NULL AUTO_INCREMENT,
    buyer_id int,
    token_id text,
    category int,
    price int NOT NULL DEFAULT '0',
    metadata int NOT NULL DEFAULT '0',
    buy_time int(11),
    start_time int(11) NOT NULL,
    end_time int(11) NOT NULL,
    created_time int(11) NOT NULL,
    PRIMARY KEY(inventory_id),
    FOREIGN KEY (buyer_id) REFERENCES users(user_id)
);

-- +goose Down
DROP TABLE inventories;
