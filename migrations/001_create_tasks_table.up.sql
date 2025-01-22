CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    priority INT NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending',
    execute_time TIMESTAMP NOT NULL,
    retry_count INT DEFAULT 0,
    max_retries INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);