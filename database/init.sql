CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(100), gender VARCHAR(100), age INTEGER);

CREATE INDEX idx_users_id ON users(id);