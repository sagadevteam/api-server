-- +goose Up
CREATE TABLE metadata (
    metadata_id int NOT NULL AUTO_INCREMENT,
    flag int NOT NULL UNIQUE,
    description text,
    PRIMARY KEY(metadata_id),
    KEY (flag)
);

-- +goose Down
DROP TABLE metadata;
