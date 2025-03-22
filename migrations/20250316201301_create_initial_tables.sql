-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS outreaches (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        address_line1 VARCHAR(255) NOT NULL,
                        address_line2 VARCHAR(255),
                        post_code VARCHAR(50) NOT NULL,
                        city VARCHAR(100) NOT NULL,
                        country VARCHAR(100) NOT NULL
    );

CREATE TABLE IF NOT EXISTS users (
                        id SERIAL PRIMARY KEY,
                        first_name VARCHAR(255) NOT NULL,
                        last_name VARCHAR(255) NOT NULL,
                        email VARCHAR(255) NOT NULL UNIQUE,
                        hashed_password VARCHAR(255) NOT NULL,
                        birthday DATE,
                        phone VARCHAR(20),
                        outreach_id INTEGER REFERENCES outreaches(id),
                        -- If a user is part of a cell, you can use a self-referential key to point to their direct leader.
                        cell_leader_id INTEGER REFERENCES users(id),
                        created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS roles (
                       id SERIAL PRIMARY KEY,
                       role_name VARCHAR(50) UNIQUE NOT NULL
);-- e.g., 'Pastor', 'Pastora', 'Primary Leader', 'Cell Leader', 'Member' ETC

CREATE TABLE IF NOT EXISTS user_roles (
                        user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                        role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
                        assigned_at TIMESTAMP DEFAULT now(),
                        PRIMARY KEY (user_id, role_id)
);

CREATE TABLE IF NOT EXISTS cell_groups (
                             id SERIAL PRIMARY KEY,
                             name VARCHAR(100),
                             leader_id INTEGER REFERENCES users(id)  -- The cell leader for this group
);

CREATE TABLE IF NOT EXISTS cell_group_members (
                                    cell_group_id INTEGER REFERENCES cell_groups(id),
                                    user_id INTEGER REFERENCES users(id),
                                    PRIMARY KEY (cell_group_id, user_id)
);

CREATE TABLE IF NOT EXISTS ministries (
                            id SERIAL PRIMARY KEY,
                            outreach_id INTEGER NOT NULL REFERENCES outreaches(id),
                            name VARCHAR(255) NOT NULL,
                            description TEXT,
                            leader VARCHAR(255) NOT NULL,
                            date_time TIMESTAMP,
                            meeting_location VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS cell_group_members;
DROP TABLE IF EXISTS cell_groups;
DROP TABLE IF EXISTS ministries;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS outreaches;
-- +goose StatementEnd
