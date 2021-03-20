-- +goose Up

CREATE TABLE votes
(
    id          SERIAL PRIMARY KEY,
    message_id  BIGINT UNSIGNED REFERENCES messages (id),
    description VARCHAR(255),
    created_at  TIMESTAMP NOT NULL DEFAULT current_timestamp
);

CREATE TABLE vote_options
(
    id      SERIAL PRIMARY KEY,
    vote_id BIGINT UNSIGNED REFERENCES votes (id),
    label   VARCHAR(50) NOT NULL
);

CREATE TABLE vote_logs
(
    id             SERIAL PRIMARY KEY,
    vote_option_id BIGINT UNSIGNED REFERENCES vote_options (id),
    user_id        INT UNSIGNED REFERENCES users (id),
    created_at     TIMESTAMP NOT NULL DEFAULT current_timestamp
);

-- +goose Down
DROP TABLE vote_logs, vote_options, votes;
