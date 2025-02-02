CREATE TABLE upload_project_schedules(
    id CHAR(36) NOT NULL,
    project_id CHAR(36) NOT NULL,
    created_at CHAR(36) NOT NULL,
    updated_at CHAR(36) NOT NULL,
    PRIMARY KEY(id)
);
ALTER TABLE
    upload_project_schedules ADD CONSTRAINT upload_project_schedules_project_id_foreign FOREIGN KEY(project_id) REFERENCES projects(id);