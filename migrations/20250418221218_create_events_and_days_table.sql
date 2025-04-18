-- +goose Up
-- +goose StatementBegin
CREATE TABLE events (
                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        title TEXT NOT NULL,
                        description TEXT,
                        additional_information TEXT,
                        location TEXT,
                        registration_link TEXT,
                        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE event_days (
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
                            date DATE NOT NULL,
                            start_time TIME,
                            end_time TIME,
                            created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS event_days;
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
