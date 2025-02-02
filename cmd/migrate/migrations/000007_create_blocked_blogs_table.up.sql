CREATE TABLE blocked_blogs(
    id CHAR(36) NOT NULL,
    blog_id CHAR(36) NOT NULL,
    reason TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE
    blocked_blogs ADD CONSTRAINT blocked_blogs_blog_id_foreign FOREIGN KEY(blog_id) REFERENCES blogs(id);