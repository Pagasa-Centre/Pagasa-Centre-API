-- +goose Up
-- +goose StatementBegin

-- Insert roles
INSERT INTO roles (role_name) VALUES
              ('Primary'),
              ('Pastor'),
              ('Media Ministry Leader'),
              ('Media Ministry Member'),
              ('Production Ministry Member'),
              ('Production Ministry Leader'),
              ('Children''s Ministry Member'),
              ('Children''s Ministry Leader'),
              ('Music Ministry Member'),
              ('Music Ministry Leader'),
              ('Creative Arts Ministry Member'),
              ('Creative Arts Ministry Leader'),
              ('Ushering & Security Ministry Leader'),
              ('Ushering & Security Ministry Member'),
              ('Disciple'),
              ('Leader'),
              ('Church Member'),
              ('Editor'),
              ('Photographer'),
              ('Sunday School Teacher'),
              ('Admin'),
              ('Monitor'),
              ('Livestreamer'),
              ('IT Ministry');

-- Insert outreach and get its ID
WITH inserted_outreach AS (
INSERT INTO outreaches (name, address_line1, address_line2, post_code, city, country)
VALUES ('Pagasa Centre Dagenham', 'Castle Green, Gale St', 'Dagenham', 'RM9 4UN', 'London', 'UK')
    RETURNING id
    )

-- Insert ministries using outreach ID
INSERT INTO ministries (outreach_id, name, description, meeting_day, start_time, meeting_location)
SELECT
    inserted_outreach.id,
    ministry.name,
    ministry.description,
    ministry.meeting_day,
    ministry.start_time,
    ministry.meeting_location
FROM inserted_outreach,
     (VALUES
          ('Production Ministry', 'The Production ministry is responsible for transporting, assembling, and storing the church''s assets and equipment.', 'Sunday', TIMESTAMP '1970-01-01 14:00:00', 'Jo Richardson Community School'),
          ('Children''s Ministry', 'Our Children''s Ministry is dedicated to nurturing their spiritual growth and helping them discover the love of Jesus Christ.', 'Sunday', TIMESTAMP '1970-01-01 15:00:00', 'Jo Richardson Community School'),
          ('Media Ministry', 'The Media Ministry is the Church''s evangelistic extension that focuses on using media to spread the word of God.', 'Sunday', TIMESTAMP '1970-01-01 14:00:00', 'Jo Richardson Community School'),
          ('Creative Arts Ministry', 'Our Creative Arts Ministry is dedicated to creating a vibrant space where these gifts can flourish, and where we can collectively use them to glorify the Lord.', 'Sunday', TIMESTAMP '1970-01-01 14:30:00', 'Jo Richardson Community School - Drama Studio 2'),
          ('Music Ministry', 'In this ministry, we understand that genuine worship goes beyond external performances.', 'Saturday', TIMESTAMP '1970-01-01 09:00:00', 'Jo Richardson Community School - Drama Studio 2'),
          ('Ushering & Security Ministry', 'The Ushers are the first representative of Jesus Christ for a worship service.', 'Sunday', TIMESTAMP '1970-01-01 14:00:00', 'Jo Richardson Community School - Drama Studio 2')
     ) AS ministry(name, description, meeting_day, start_time, meeting_location);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Delete ministries by name (assuming they were all inserted under the same outreach)
DELETE FROM ministries
WHERE name IN (
               'Production Ministry',
               'Children''s Ministry',
               'Media Ministry',
               'Creative Arts Ministry',
               'Music Ministry',
               'Ushering & Security'
    );

-- Remove dependent user_roles entries
DELETE FROM user_roles
WHERE role_id IN (
    SELECT id FROM roles
    WHERE role_name IN (
                        'Primary',
                        'Pastor',
                        'Ministry Leader',
                        'Disciple',
                        'Leader',
                        'Church Member',
                        'Media Ministry',
                        'Production Ministry',
                        'Children''s Ministry',
                        'Music Ministry',
                        'Creative Arts Ministry',
                        'Ushering & Security Ministry',
                        'Editor',
                        'Photographer',
                        'Sunday School Teacher',
                        'Admin',
                        'Monitor',
                        'Livestreamer',
                        'IT Ministry'
        )
);

-- Delete roles
DELETE FROM roles
WHERE role_name IN (
                    'Primary',
                    'Pastor',
                    'Ministry Leader',
                    'Disciple',
                    'Leader',
                    'Church Member',
                    'Media Ministry',
                    'Production Ministry',
                    'Children''s Ministry',
                    'Music Ministry',
                    'Creative Arts Ministry',
                    'Ushering & Security Ministry',
                    'Editor',
                    'Photographer',
                    'Sunday School Teacher',
                    'Admin',
                    'Monitor',
                    'Livestreamer',
                    'IT Ministry'
    );

-- Remove outreach
DELETE FROM outreaches
WHERE name = 'Pagasa Centre Dagenham'
  AND post_code = 'RM9 4UN'
  AND city = 'London'
  AND country = 'England';

-- +goose StatementEnd