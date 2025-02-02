CREATE TABLE blocked_projects(
    id CHAR(36) NOT NULL,
    project_id CHAR(36) NOT NULL,
    reason TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE
    blocked_projects ADD CONSTRAINT blocked_projects_project_id_foreign FOREIGN KEY(id) REFERENCES projects(id);