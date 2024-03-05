-- +goose Up

CREATE TABLE urls (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    url TEXT NOT NULL,
    short_url UUID NOT NULL DEFAULT uuid_generate_v4(),
    clicks INTEGER DEFAULT 0
);

-- +goose Down
DROP TABLE urls;