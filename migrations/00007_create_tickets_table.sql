-- +goose Up
CREATE TABLE tickets (
    ticket_id int NOT NULL AUTO_INCREMENT,
    inventory_id int,
    user_id int,
    token_id int(11),
    time int(11),
    PRIMARY KEY(ticket_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (inventory_id) REFERENCES inventories(inventory_id),
    KEY (time, token_id)
);

-- +goose Down
DROP TABLE tickets;