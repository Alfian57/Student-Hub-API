CREATE TABLE projects(
    id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    category_id CHAR(36) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(255) NULL,
    thumbnail VARCHAR(255) NOT NULL,
    is_publish BOOLEAN NOT NULL,
    code_link VARCHAR(255) NULL,
    app_link VARCHAR(255) NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE
    projects ADD CONSTRAINT projects_slug_unique UNIQUE(slug);
ALTER TABLE
    projects ADD CONSTRAINT projects_user_id_foreign FOREIGN KEY(user_id) REFERENCES users(id);
ALTER TABLE
    projects ADD CONSTRAINT projects_category_id_foreign FOREIGN KEY(category_id) REFERENCES categories(id);