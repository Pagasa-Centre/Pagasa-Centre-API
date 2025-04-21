-- +goose Up
-- +goose StatementBegin

-- 1. Drop leader_id column
ALTER TABLE ministries
DROP COLUMN IF EXISTS leader_id;

-- 2. Add short_description, long_description, thumbnail_url
ALTER TABLE ministries
    ADD COLUMN short_description VARCHAR(180),
ADD COLUMN long_description TEXT,
ADD COLUMN thumbnail_url TEXT;

-- 3. Create ministry_activities table
CREATE TABLE IF NOT EXISTS ministry_activities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ministry_id UUID NOT NULL REFERENCES ministries(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- 4. Insert activities for each ministry

-- MUSIC MINISTRY
INSERT INTO ministry_activities (ministry_id, name)
SELECT id, activity FROM ministries,
                         (VALUES
                              ('Vocal Training'),
                              ('Musical Instruments'),
                              ('Sound Engineering'),
                              ('Worship Leading')
                         ) AS a(activity)
WHERE ministries.name = 'Music Ministry';

-- MEDIA MINISTRY
INSERT INTO ministry_activities (ministry_id, name)
SELECT id, activity FROM ministries,
                         (VALUES
                              ('Photography'),
                              ('Videography'),
                              ('Editing'),
                              ('Projection'),
                              ('Live Streaming')
                         ) AS a(activity)
WHERE ministries.name = 'Media Ministry';

-- PRODUCTION MINISTRY
INSERT INTO ministry_activities (ministry_id, name)
SELECT id, activity FROM ministries,
                         (VALUES
                              ('Equipment Setup'),
                              ('Sound System Management'),
                              ('Lighting Control'),
                              ('Stage Management')
                         ) AS a(activity)
WHERE ministries.name = 'Production Ministry';

-- CHILDREN'S MINISTRY
INSERT INTO ministry_activities (ministry_id, name)
SELECT id, activity FROM ministries,
                         (VALUES
                              ('Bible Stories'),
                              ('Worship Songs'),
                              ('Arts and Crafts'),
                              ('Group Activities'),
                              ('Lesson Teaching')
                         ) AS a(activity)
WHERE ministries.name = 'Children''s Ministry';

-- CREATIVE ARTS MINISTRY (with made-up activities)
INSERT INTO ministry_activities (ministry_id, name)
SELECT id, activity FROM ministries,
                         (VALUES
                              ('Spoken Word Practice'),
                              ('Drama Rehearsals'),
                              ('Choreographed Worship'),
                              ('Stage Presence Workshops')
                         ) AS a(activity)
WHERE ministries.name = 'Creative Arts Ministry';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop ministry_activities table
DROP TABLE IF EXISTS ministry_activities;

-- Remove added columns
ALTER TABLE ministries
DROP COLUMN IF EXISTS short_description,
DROP COLUMN IF EXISTS long_description,
DROP COLUMN IF EXISTS thumbnail_url;

-- Restore leader_id (NOTE: nullable since we no longer seed it)
ALTER TABLE ministries
    ADD COLUMN leader_id UUID REFERENCES users(id);

-- +goose StatementEnd