-- +goose Up

CREATE TABLE attachment_types
(
    id    INT UNSIGNED PRIMARY KEY,
    alias VARCHAR(64) NOT NULL,
    title VARCHAR(64) NOT NULL
);
INSERT INTO attachment_types (id, alias, title)
VALUES (1, 'photo', 'Photos'),
       (2, 'video', 'Videos'),
       (3, 'link', 'Links');

CREATE TABLE attachments
(
    id                 BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    message_id         BIGINT UNSIGNED REFERENCES messages (id),
    attachment_type_id INT UNSIGNED REFERENCES attachment_types (id),
    content_url        VARCHAR(2048) NOT NULL,
    created_at         TIMESTAMP     NOT NULL DEFAULT current_timestamp,
    INDEX (message_id)
);

-- +goose Down
DROP TABLE attachments, attachment_types;
