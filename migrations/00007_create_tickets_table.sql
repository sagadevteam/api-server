-- +goose Up
CREATE TABLE tickets (
    ticket_id int NOT NULL AUTO_INCREMENT,
    inventory_id int NOT NULL,
    user_id int,
    time int(11) NOT NULL,
    PRIMARY KEY(ticket_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (inventory_id) REFERENCES inventories(inventory_id),
    KEY (time)
);

-- +goose Down
DROP TABLE tickets;