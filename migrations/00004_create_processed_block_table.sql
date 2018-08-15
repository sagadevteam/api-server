-- +goose Up
CREATE TABLE processed_block (
    id int NOT NULL AUTO_INCREMENT,
    block_height int NOT NULL DEFAULT '0',
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE processed_block;
