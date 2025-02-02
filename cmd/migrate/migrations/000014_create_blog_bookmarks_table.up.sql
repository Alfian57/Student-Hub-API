CREATE TABLE blog_bookmarks(
    id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    blog_id CHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE
    blog_bookmarks ADD CONSTRAINT blog_bookmarks_blog_id_foreign FOREIGN KEY(blog_id) REFERENCES blogs(id);