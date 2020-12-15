-- +goose Up

ALTER TABLE users
    ADD phone      VARCHAR(50) UNIQUE,
    ADD password   VARCHAR(60) NOT NULL,
    ADD bio        VARCHAR(255)  DEFAULT NULL,
    ADD photo_url   VARCHAR(1024) DEFAULT NULL,
    ADD created_at TIMESTAMP   NOT NULL DEFAULT current_timestamp,
    ADD updated_at TIMESTAMP   NOT NULL DEFAULT current_timestamp;
ALTER TABLE users
    ADD INDEX phone_idx (phone);
ALTER TABLE users
    ADD INDEX name_idx (name);

-- +goose Down
ALTER TABLE users
    DROP INDEX phone_idx;
ALTER TABLE users
    DROP INDEX name_idx;
ALTER TABLE users
    DROP phone,
    DROP password,
    DROP bio,
    DROP photo_url,
    DROP created_at,
    DROP updated_at;
