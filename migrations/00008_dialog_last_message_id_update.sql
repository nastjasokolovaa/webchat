-- +goose Up
-- +goose StatementBegin
CREATE TRIGGER dialog_last_message_id_update
    AFTER INSERT ON messages
    FOR EACH ROW
    BEGIN
        UPDATE dialogs
        SET last_message_id = NEW.id
        WHERE id = NEW.dialog_id;
    END;
-- +goose StatementEnd

-- +goose Down

DROP TRIGGER dialog_last_message_id_update;
