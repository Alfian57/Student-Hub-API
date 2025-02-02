CREATE TABLE project_ratings(
    id CHAR(36) NOT NULL,
    project_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    rating SMALLINT NOT NULL CHECK (rating >= 0 AND rating <= 5),
    review TEXT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE
    project_ratings ADD CONSTRAINT project_ratings_project_id_foreign FOREIGN KEY(project_id) REFERENCES projects(id);