-- +goose Up
CREATE TABLE IF NOT EXISTS role_permissions(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    role_id BIGINT UNSIGNED NOT NULL,
    permission_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);

-- Admin (role_id=1) gets all permissions
INSERT INTO role_permissions(role_id, permission_id)
SELECT 1, id FROM permissions;

-- Regular user (role_id=2) gets only user:read
INSERT INTO role_permissions(role_id, permission_id)
SELECT 2, id FROM permissions WHERE name IN ('user:read');

-- +goose Down
DROP TABLE IF EXISTS role_permissions;