DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'category_type') THEN
        CREATE TYPE category_type AS ENUM ('project', 'blog');
    END IF;
END $$;

CREATE TABLE categories(
    id CHAR(36) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    type category_type NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE
    categories ADD CONSTRAINT categories_slug_unique UNIQUE(slug);