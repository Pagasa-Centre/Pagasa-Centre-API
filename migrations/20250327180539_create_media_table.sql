-- +goose Up
CREATE TABLE IF NOT EXISTS media (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT,
        youtube_video_id VARCHAR(100) NOT NULL UNIQUE,
        category VARCHAR(50) NOT NULL, -- 'preachings', 'bible_study', 'evangelistic_night'
        published_at TIMESTAMP NOT NULL,
        thumbnail_url TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT now(),
        updated_at TIMESTAMP DEFAULT now()
    );

-- +goose Down
DROP TABLE IF EXISTS media;