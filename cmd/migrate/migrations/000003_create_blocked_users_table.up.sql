CREATE TABLE blocked_users(
    id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    reason TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE
    blocked_users ADD CONSTRAINT blocked_users_user_id_foreign FOREIGN KEY(user_id) REFERENCES users(id);