-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS games
(
    game_id  uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        text,
    description text,
    price       numeric,
    genre       text,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE games;
-- +goose StatementEnd
