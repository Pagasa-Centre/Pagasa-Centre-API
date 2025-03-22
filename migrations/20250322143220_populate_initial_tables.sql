-- +goose Up
-- +goose StatementBegin
INSERT INTO roles (role_name) VALUES
                                  ('Primary'),
                                  ('Pastor'),
                                  ('Ministry Leader'),
                                  ('Disciple'),
                                  ('Leader'),
                                  ('Church Member'),
                                  ('Media Ministry'),
                                  ('Production Ministry'),
                                  ('Children''s Ministry'),
                                  ('Music Ministry'),
                                  ('Creative Arts Ministry'),
                                  ('Ushering & Security Ministry'),
                                  ('Editor'),
                                  ('Photographer'),
                                  ('Sunday School Teacher'),
                                  ('Admin'),
                                  ('Monitor'),
                                  ('Livestreamer'),
                                  ('IT Ministry');

INSERT INTO outreaches (name, address_line1, address_line2, post_code, city, country)
VALUES (
           'Pagasa Centre Dagenham',
           'Castle Green, Gale St',
           'Dagenham',
           'RM9 4UN',
           'London',
           'England'
       );

INSERT INTO ministries (outreach_id, name, description, meeting_day, meeting_time, meeting_location)
VALUES
    (
        1,
        'Production Ministry',
        'The Production ministry is responsible for transporting, assembling, and storing the church''s assets and equipment.',
        'Sunday',
        '1970-01-01 14:00:00'::timestamp,
        'Jo Richardson Community School'
    ),
    (
        1,
        'Children''s Ministry',
        'Our Children''s Ministry is dedicated to nurturing their spiritual growth and helping them discover the love of Jesus Christ.',
        'Sunday',
        '1970-01-01 15:00:00'::timestamp,
        'Jo Richardson Community School'
    ),
    (
        1,
        'Media Ministry',
        'The Media Ministry is the Church''s evangelistic extension that focuses on using media to spread the word of God.',
        'Sunday',
        '1970-01-01 14:00:00'::timestamp,
        'Jo Richardson Community School'
    ),
    (
        1,
        'Creative Arts Ministry',
        'Our Creative Arts Ministry is dedicated to creating a vibrant space where these gifts can flourish, and where we can collectively use them to glorify the Lord.',
        'Sunday',
        '1970-01-01 14:30:00'::timestamp,
        'Jo Richardson Community School - Drama Studio 2'
    ),
    (
        1,
        'Music Ministry',
        'In this ministry, we understand that genuine worship goes beyond external performances.',
        'Saturday',
        '1970-01-01 09:00:00'::timestamp,
        'Jo Richardson Community School - Drama Studio 2'
    ),
    (
        1,
        'Ushering & Security',
        'The Ushers are the first representative of Jesus Christ for a worship service.',
        'Sunday',
        '1970-01-01 14:00:00'::timestamp,
        'Jo Richardson Community School - Drama Studio 2'
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- First, remove the inserted ministries.
DELETE FROM ministries
WHERE outreach_id = 1
  AND name IN (
               'Production Ministry',
               'Children''s Ministry',
               'Media Ministry',
               'Creative Arts Ministry',
               'Music Ministry',
               'Ushering & Security'
    );

-- Then, remove dependent user_roles entries.
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

-- Now, delete the roles.
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

-- Update users to remove the reference to the outreach.
UPDATE users
SET outreach_id = NULL
WHERE outreach_id = (
    SELECT id FROM outreaches
    WHERE name = 'Pagasa Centre Dagenham'
      AND post_code = 'RM9 4UN'
      AND city = 'London'
      AND country = 'England'
);

-- Delete the outreach record.
DELETE FROM outreaches
WHERE name = 'Pagasa Centre Dagenham'
  AND post_code = 'RM9 4UN'
  AND city = 'London'
  AND country = 'England';
-- +goose StatementEnd