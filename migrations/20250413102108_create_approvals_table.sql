-- +goose Up
CREATE TABLE approvals (
            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
            requester_id UUID NOT NULL,
            updated_by UUID REFERENCES users(id),
            type TEXT NOT NULL,
            requested_role TEXT NOT NULL,
            reason TEXT,
            status TEXT NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE approvals;