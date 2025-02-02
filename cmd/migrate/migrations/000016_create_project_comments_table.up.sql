CREATE TABLE project_comments(
    id CHAR(36) NOT NULL,
    project_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    comment TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE
    project_comments ADD CONSTRAINT project_comments_user_id_foreign FOREIGN KEY(user_id) REFERENCES users(id);
ALTER TABLE
    project_comments ADD CONSTRAINT project_comments_project_id_foreign FOREIGN KEY(project_id) REFERENCES projects(id);