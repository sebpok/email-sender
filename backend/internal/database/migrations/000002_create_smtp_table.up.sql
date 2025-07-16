-- migrations/000002_create_smtp_table.up.sql
CREATE TABLE smtp (
    id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE NOT NULL,
    host VARCHAR(255) NOT NULL,
    port INTEGER NOT NULL,
    username VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    encryption VARCHAR(10) NOT NULL, -- e.g. 'ssl', 'tls', 'starttls', 'none'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_users
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_smtp_user_id ON smtp(user_id);