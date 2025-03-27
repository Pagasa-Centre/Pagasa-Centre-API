-- +goose Up
-- +goose StatementBegin

UPDATE outreaches
SET thumbnail_url = 'https://cdn.prod.website-files.com/6469d767492ea69c34c8827d/646a8f02f6f8fe6b1f83c58b_jo%20richardson.jpg'
WHERE name = 'Pagasa Centre Dagenham';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

UPDATE outreaches
SET thumbnail_url = NULL
WHERE name = 'Pagasa Centre Dagenham';

-- +goose StatementEnd