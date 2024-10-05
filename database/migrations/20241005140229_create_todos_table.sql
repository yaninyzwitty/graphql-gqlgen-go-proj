-- +goose Up
-- +goose StatementBegin
CREATE TABLE todos (
    id UUID PRIMARY KEY,
    text TEXT NOT NULL,
    done BOOLEAN DEFAULT false,
    user_id UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todos;
-- +goose StatementEnd


