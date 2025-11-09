-- +goose Up
-- +goose StatementBegin
GRANT ALL PRIVILEGES ON DATABASE e_shop TO app_user;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
REVOKE ALL PRIVILEGES ON DATABASE e_shop FROM app_user;
-- +goose StatementEnd
