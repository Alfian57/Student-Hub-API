CREATE TABLE collaborators(
    id CHAR(36) NOT NULL,
    project_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    role VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE
    collaborators ADD CONSTRAINT collaborators_user_id_foreign FOREIGN KEY(user_id) REFERENCES users(id);
ALTER TABLE
    collaborators ADD CONSTRAINT collaborators_project_id_foreign FOREIGN KEY(project_id) REFERENCES projects(id);