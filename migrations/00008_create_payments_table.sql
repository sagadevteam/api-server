-- +goose Up
CREATE TABLE payments (
    payment_id int NOT NULL AUTO_INCREMENT,
    ticket_id int NOT NULL,
    user_id int NOT NULL,
    created_time int(11) NOT NULL,
    deleted_time int(11),
    PRIMARY KEY(payment_id),
    FOREIGN KEY (ticket_id) REFERENCES tickets(ticket_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    KEY (created_time, deleted_time)
);

-- +goose Down
DROP TABLE payments;