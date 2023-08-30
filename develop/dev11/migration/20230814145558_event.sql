-- +goose Up
-- +goose StatementBegin
CREATE TABLE event
(
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER NOT NULL,
    title       TEXT    NOT NULL,
    description TEXT    NOT NULL DEFAULT '',
    date        TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE event;
-- +goose StatementEnd
