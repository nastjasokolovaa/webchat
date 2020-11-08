-- +goose Up

CREATE TABLE dialogs
(
    id       BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    owner_id INT UNSIGNED REFERENCES users (id)
);

CREATE TABLE messages
(
    id         BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    dialog_id  BIGINT UNSIGNED REFERENCES dialogs (id),
    sender_id  INT UNSIGNED REFERENCES users (id),
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    message    TEXT      NOT NULL,
    INDEX (dialog_id)
);

CREATE TABLE dialog_users
(
    id                   BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    dialog_id            BIGINT UNSIGNED REFERENCES dialogs (id),
    user_id              INT UNSIGNED REFERENCES users (id),
    last_read_message_id BIGINT UNSIGNED REFERENCES messages (id),
    UNIQUE INDEX (dialog_id, user_id),
    INDEX (user_id)
);

-- +goose Down
DROP TABLE dialog_users, messages, dialogs;
