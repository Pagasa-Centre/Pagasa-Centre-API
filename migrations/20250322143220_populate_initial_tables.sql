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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
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

DELETE FROM outreaches
WHERE name = 'Pagasa Centre Dagenham'
  AND post_code = 'RM9 4UN'
  AND city = 'London'
  AND country = 'England';
-- +goose StatementEnd
