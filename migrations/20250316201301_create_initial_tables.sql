-- +goose Up
-- +goose StatementBegin

-- Enable pgcrypto for UUID generation
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS outreaches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    address_line1 VARCHAR(255) NOT NULL,
    address_line2 VARCHAR(255),
    post_code VARCHAR(50),
    city VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL
    );

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    hashed_password VARCHAR(255) NOT NULL,
    birthday DATE,
    phone VARCHAR(20),
    outreach_id UUID REFERENCES outreaches(id),
    cell_leader_id UUID REFERENCES users(id), -- self-referential foreign key
    created_at TIMESTAMP DEFAULT now()
    );

CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role_name VARCHAR(50) UNIQUE NOT NULL
    );

CREATE TABLE IF NOT EXISTS user_roles (
     user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
    assigned_at TIMESTAMP DEFAULT now(),
    PRIMARY KEY (user_id, role_id)
    );

CREATE TABLE IF NOT EXISTS cell_groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100),
    leader_id UUID REFERENCES users(id)
    );

CREATE TABLE IF NOT EXISTS cell_group_members (
    cell_group_id UUID REFERENCES cell_groups(id),
    user_id UUID REFERENCES users(id),
    PRIMARY KEY (cell_group_id, user_id)
    );

CREATE TABLE IF NOT EXISTS ministries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    outreach_id UUID NOT NULL REFERENCES outreaches(id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    leader_id UUID REFERENCES users(id),
    meeting_day VARCHAR(15),
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
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