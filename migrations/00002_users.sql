-- +goose Up
CREATE TABLE users
(
    id   INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL
);

-- +goose Down
DROP TABLE users;
