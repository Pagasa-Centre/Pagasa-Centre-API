-- +goose Up
-- +goose StatementBegin

-- Add optional columns if they don't exist yet
ALTER TABLE outreaches
    ADD COLUMN IF NOT EXISTS region VARCHAR(100),
    ADD COLUMN IF NOT EXISTS venue_name VARCHAR(255);

-- Insert new outreaches
INSERT INTO outreaches (name, address_line1, address_line2, post_code, city, country, region, venue_name)
VALUES
    (
        'Pagasa Centre Bray',
        'Taylor Centre, Vevay Road',
        'Bray Co. Wicklow',
        'A98 E220',
        'Bray',
        'Ireland',
        'Bray',
        'Taylor Centre'
    ),
    (
        'Pagasa Centre Pampanga',
        '2nd Floor of Juliez Manukan Blds',
        'San Matias Highway, Santo Tomas',
        NULL,
        'Santo Tomas',
        'Philippines',
        'Pampanga',
        NULL
    ),
    (
        'Pagasa Centre Bedfordshire',
        '30 Bunyan Road',
        'Kempston',
        'MK42 8HL',
        'Bedfordshire',
        'UK',
        'Bedfordshire',
        NULL
    ),
    (
        'Pagasa Centre Reading',
        'Neville Hall, Milley Rd',
        'Waltham St Lawrence',
        'RG10 0JP',
        'Reading',
        'UK',
        'Reading',
        'Neville Hall'
    ),
    (
        'Pagasa Centre Harwich',
        'Mayflower Primary School Hall, Main Road, Dovercourt',
        'Harwich, Essex',
        'CO12 4AJ',
        'Harwich',
        'UK',
        'Harwich',
        'Mayflower Primary School Hall'
    ),
    (
        'Pagasa Centre Stratford Upon Avon',
        'Ken Kennett Centre',
        '100 Justins Avenue',
        'CV37 0DA',
        'Stratford Upon Avon',
        'UK',
        'Stratford Upon Avon',
        'Ken Kennett Centre'
    ),
    (
        'Pagasa Centre Banga',
        'Ruiz Compound',
        'Bgy Kusan, Barrio 8',
        NULL,
        'Banga',
        'Philippines',
        'South Cotabato',
        NULL
    ),
    (
        'Pagasa Centre West Midlands & Worcestershire',
        'The Old Library Centre, 65 Ombersley Street East',
        'Droitwich Spa',
        'WR9 8QS',
        'Droitwich',
        'UK',
        'West Midlands & Worcestershire',
        'The Old Library Centre'
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM outreaches WHERE name IN (
                                      'Pagasa Centre Bray',
                                      'Pagasa Centre Pampanga',
                                      'Pagasa Centre Bedfordshire',
                                      'Pagasa Centre Reading',
                                      'Pagasa Centre Harwich',
                                      'Pagasa Centre Stratford Upon Avon',
                                      'Pagasa Centre Banga',
                                      'Pagasa Centre West Midlands & Worcestershire'
    );
-- +goose StatementEnd