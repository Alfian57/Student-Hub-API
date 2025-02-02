CREATE TABLE blog_comments(
    id CHAR(36) NOT NULL,
    blog_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    comment TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE
    blog_comments ADD CONSTRAINT blog_comments_user_id_foreign FOREIGN KEY(user_id) REFERENCES users(id);
ALTER TABLE
    blog_comments ADD CONSTRAINT blog_comments_blog_id_foreign FOREIGN KEY(blog_id) REFERENCES blogs(id);