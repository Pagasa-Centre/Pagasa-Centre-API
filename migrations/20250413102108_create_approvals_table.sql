-- +goose Up
CREATE TABLE approvals (
                           id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                           requester_id UUID NOT NULL,
                           approver_id UUID NOT NULL,
                           type TEXT NOT NULL,
                           status TEXT NOT NULL,
                           created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE approvals;