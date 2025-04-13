-- +goose Up
-- +goose StatementBegin

CREATE TABLE blogs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    body TEXT NOT NULL
) STRICT;

-- CREATE TABLE tags ();
--
-- CREATE TABLE projects ();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE blogs;
-- +goose StatementEnd
