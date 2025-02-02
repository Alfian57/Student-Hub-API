CREATE TABLE blogs(
    id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    category_id CHAR(36) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    thumbnail VARCHAR(255) NULL,
    is_publish BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE
    blogs ADD CONSTRAINT blogs_slug_unique UNIQUE(slug);
ALTER TABLE
    blogs ADD CONSTRAINT blogs_category_id_foreign FOREIGN KEY(category_id) REFERENCES categories(id);