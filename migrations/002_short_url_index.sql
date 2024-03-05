-- +goose Up

CREATE INDEX short_url_index ON urls(short_url);

-- +goose Down
DROP INDEX short_url_index;