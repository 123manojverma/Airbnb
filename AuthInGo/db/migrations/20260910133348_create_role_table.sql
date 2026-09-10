-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO roles(name,description) VALUES
('admin','Administrator with full access'),
('user','Regular user with limited access'),
('moderator','Moderator with elevated privileges');

-- +goose Down
DROP TABLE IF EXISTS roles;