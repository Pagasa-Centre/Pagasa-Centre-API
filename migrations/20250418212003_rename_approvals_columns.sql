-- +goose Up
ALTER TABLE approvals
    RENAME COLUMN type TO requested_role;

ALTER TABLE approvals
    ADD COLUMN type TEXT;

-- +goose Down
ALTER TABLE approvals
DROP COLUMN type;

ALTER TABLE approvals
    RENAME COLUMN requested_role TO type;