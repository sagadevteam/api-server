-- +goose Up
CREATE TABLE users (
    id int NOT NULL,
    email text,
    password text,
    eth_address text,
    is_admin tinyint(1),
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE users;
