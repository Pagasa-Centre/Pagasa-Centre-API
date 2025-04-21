-- +goose Up
-- +goose StatementBegin
ALTER TABLE approvals
    ADD COLUMN ministry_id UUID REFERENCES ministries(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE approvals
DROP COLUMN ministry_id;
-- +goose StatementEnd
