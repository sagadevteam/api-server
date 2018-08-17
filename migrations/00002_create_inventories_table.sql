-- +goose Up
CREATE TABLE inventories (
    inventory_id int NOT NULL AUTO_INCREMENT,
    price int NOT NULL DEFAULT '0',
    metadata int NOT NULL DEFAULT '0',
    start_time int(11) NOT NULL,
    end_time int(11) NOT NULL,
    created_time int(11) NOT NULL,
    PRIMARY KEY(inventory_id),
    KEY (price, metadata, start_time, end_time, created_time)
);

-- +goose Down
DROP TABLE inventories;
