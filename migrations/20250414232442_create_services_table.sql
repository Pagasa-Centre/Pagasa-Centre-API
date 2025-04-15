-- +goose Up
CREATE TABLE outreach_services (
       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
       day VARCHAR(50) NOT NULL,
       start_time TIME NOT NULL,
       end_time TIME NOT NULL,
       outreach_id UUID NOT NULL,
       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
       updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
       CONSTRAINT fk_outreach
           FOREIGN KEY (outreach_id)
               REFERENCES outreaches(id)
               ON DELETE CASCADE
);

INSERT INTO outreach_services (day, start_time, end_time, outreach_id)
SELECT
    'Sunday',
    '14:00',
    '16:30',
    id
FROM outreaches
WHERE name = 'Pagasa Centre Dagenham';

-- +goose Down
DELETE FROM outreach_services
WHERE outreach_id = (
    SELECT id FROM outreaches WHERE name = 'Pagasa Centre Dagenham'
)
          AND day = 'Sunday'
  AND start_time = '14:00'
  AND end_time = '16:30';
DROP TABLE IF EXISTS outreach_services;