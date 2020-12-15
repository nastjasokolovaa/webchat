-- +goose Up

ALTER TABLE dialogs
    ADD last_message_id BIGINT UNSIGNED;
-- +goose Down
ALTER TABLE dialogs
    DROP last_message_id;
