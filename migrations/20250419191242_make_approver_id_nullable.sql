-- +goose Up
-- +goose StatementBegin
ALTER TABLE approvals
    ALTER COLUMN approver_id DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE approvals
    ALTER COLUMN approver_id SET NOT NULL;
-- +goose StatementEnd