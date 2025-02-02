CREATE TABLE project_bookmarks(
    id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    project_id CHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE
    project_bookmarks ADD CONSTRAINT project_bookmarks_project_id_foreign FOREIGN KEY(project_id) REFERENCES projects(id);