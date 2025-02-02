CREATE TABLE donations(
    id CHAR(36) NOT NULL,
    project_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    amount BIGINT NOT NULL CHECK (amount >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE
    donations ADD CONSTRAINT donations_user_id_foreign FOREIGN KEY(user_id) REFERENCES users(id);
ALTER TABLE
    donations ADD CONSTRAINT donations_project_id_foreign FOREIGN KEY(project_id) REFERENCES projects(id);