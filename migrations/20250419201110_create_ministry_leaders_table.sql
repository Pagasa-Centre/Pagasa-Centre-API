-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ministry_leaders (
      user_id UUID NOT NULL,
      ministry_id UUID NOT NULL,
      assigned_at TIMESTAMP DEFAULT now(),

      PRIMARY KEY (user_id, ministry_id),

      FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
      FOREIGN KEY (ministry_id) REFERENCES ministries(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ministry_leaders;
-- +goose StatementEnd
