-- +goose Up
CREATE TABLE users (
    user_id int NOT NULL AUTO_INCREMENT,
    email text NOT NULL,
    password text NOT NULL,
    eth_addr text NOT NULL,
    eth_value text NOT NULL,
    saga_point text NOT NULL,
    is_admin tinyint(1) DEFAULT '0',
    PRIMARY KEY(user_id)
);

-- +goose Down
DROP TABLE users;
