-- +goose Up
CREATE TABLE users (
    user_id int NOT NULL AUTO_INCREMENT,
    email varchar(512) NOT NULL,
    password varchar(100) NOT NULL,
    eth_addr varchar(60) NOT NULL,
    eth_priv varchar(200) NOT NULL,
    eth_value text NOT NULL,
    saga_point int(11) NOT NULL,
    is_admin tinyint(1) DEFAULT '0',
    PRIMARY KEY(user_id),
    UNIQUE (email)
);

-- +goose Down
DROP TABLE users;
