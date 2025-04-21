-- +goose Up
-- +goose StatementBegin

-- Add optional columns if they don't exist yet
ALTER TABLE outreaches
    ADD COLUMN IF NOT EXISTS region VARCHAR(100),
    ADD COLUMN IF NOT EXISTS venue_name VARCHAR(255),
    ADD COLUMN IF NOT EXISTS thumbnail_url TEXT;


-- Insert new outreaches
INSERT INTO outreaches (name, address_line1, address_line2, post_code, city, country, region, venue_name,thumbnail_url)
VALUES
    (
        'Pagasa Centre Bray',
        'Taylor Centre, Vevay Road',
        'Bray Co. Wicklow',
        'A98 E220',
        'Bray',
        'Ireland',
        'Bray',
        'Taylor Centre',
     'https://cdn.prod.website-files.com/6469d767492ea69c34c8827d/64e52429d6a96262b5c2bc7e_ireland%20temp.jpeg'
    ),
    (
        'Pagasa Centre Pampanga',
        '2nd Floor of Juliez Manukan Blds',
        'San Matias Highway, Santo Tomas',
        NULL,
        'Santo Tomas',
        'Philippines',
        'Pampanga',
        NULL,
     'https://cdn.prod.website-files.com/6469d767492ea69c34c8827d/64e523c80c75eda966b8247b_filil%20church.jpg'
    ),
    (
        'Pagasa Centre Bedfordshire',
        '30 Bunyan Road',
        'Kempston',
        'MK42 8HL',
        'Bedfordshire',
        'UK',
        'Bedfordshire',
        NULL,
     'https://cdn.prod.website-files.com/6469d767492ea69c34c8827d/64e526e93be191fdc4a682a7_%E3%82%B9%E3%82%AF%E3%83%AA%E3%83%BC%E3%83%B3%E3%82%B7%E3%83%A7%E3%83%83%E3%83%88%202023-08-22%20221934.jpg'

    ),
    (
        'Pagasa Centre Reading',
        'Neville Hall, Milley Rd',
        'Waltham St Lawrence',
        'RG10 0JP',
        'Reading',
        'UK',
        'Reading',
        'Neville Hall',
        'https://cdn.prod.website-files.com/6469d767492ea69c34c8827d/654039ec2ed2e7273b6e22e1_reading.jpg'
    ),
    (
        'Pagasa Centre Harwich',
        'Mayflower Primary School Hall, Main Road, Dovercourt',
        'Harwich, Essex',
        'CO12 4AJ',
        'Harwich',
        'UK',
        'Harwich',
        'Mayflower Primary School Hall',
     'https://cdn.prod.website-files.com/6469d767492ea69c34c8827d/6609e540b03e95120f827d4e_WhatsApp%20Image%202024-03-30%20at%205.41.59%20PM.jpeg'
    ),
    (
        'Pagasa Centre Stratford Upon Avon',
        'Ken Kennett Centre',
        '100 Justins Avenue',
        'CV37 0DA',
        'Stratford Upon Avon',
        'UK',
        'Stratford Upon Avon',
        'Ken Kennett Centre',
     'https://cdn.prod.website-files.com/6469d767492ea69c34c8827d/660315e82e2e2eb16ed51714_image.png'
    ),
    (
        'Pagasa Centre Banga',
        'Ruiz Compound',
        'Bgy Kusan, Barrio 8',
        NULL,
        'Banga',
        'Philippines',
        'South Cotabato',
        NULL,
        'https://cdn.prod.website-files.com/6469d767492ea69c34c8827d/6601e720dcdada916755f7d0_banga%20church.jpeg'
    ),
    (
        'Pagasa Centre West Midlands & Worcestershire',
        'The Old Library Centre, 65 Ombersley Street East',
        'Droitwich Spa',
        'WR9 8QS',
        'Droitwich',
        'UK',
        'West Midlands & Worcestershire',
        'The Old Library Centre',
     'https://cdn.prod.website-files.com/6469d767492ea69c34c8827d/67868af10e281f8fd30f715d_image.png'
    ),
    (
        'Pagasa Centre Southend-on-sea',
        'The Cornerstone URC Church, Bournemouth Park Road',
        'Southend-on-sea, Essex',
        'SS2 5JL',
        'Southend-on-sea',
        'UK',
        'Southend-on-sea',
        'The Cornerstone URC Church',
        'https://cdn.prod.website-files.com/6469d767492ea69c34c8827d/67e5a044384ad46614e599ee_south.png'
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users
WHERE outreach_id IN (
    SELECT id FROM outreaches WHERE name IN (
                                             'Pagasa Centre Bray',
                                             'Pagasa Centre Pampanga',
                                             'Pagasa Centre Bedfordshire',
                                             'Pagasa Centre Reading',
                                             'Pagasa Centre Harwich',
                                             'Pagasa Centre Stratford Upon Avon',
                                             'Pagasa Centre Banga',
                                             'Pagasa Centre West Midlands & Worcestershire',
                                             'Pagasa Centre Southend-on-sea'
        )
);

DELETE FROM outreaches WHERE name IN (
                                      'Pagasa Centre Bray',
                                      'Pagasa Centre Pampanga',
                                      'Pagasa Centre Bedfordshire',
                                      'Pagasa Centre Reading',
                                      'Pagasa Centre Harwich',
                                      'Pagasa Centre Stratford Upon Avon',
                                      'Pagasa Centre Banga',
                                      'Pagasa Centre West Midlands & Worcestershire',
                                      'Pagasa Centre Southend-on-sea'
    );
-- +goose StatementEnd