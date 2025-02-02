package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BlogStore struct {
	db *sqlx.DB
}

type Blog struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	CategoryID uuid.UUID `json:"category_id"`
	Slug       string    `json:"slug"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Thumbnail  string    `json:"thumbnail"`
	IsPublish  string    `json:"is_publish"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type BlockedBlog struct {
	ID        uuid.UUID `json:"id"`
	BlogID    uuid.UUID `json:"blog_id"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *BlogStore) GetAll(ctx context.Context, queryParam BlogQueryParam) (*[]Blog, error) {
	query := fmt.Sprintf(`
		SELECT id, user_id, category_id, slug, title, content, thumbnail, is_publish, created_at, updated_at
		FROM blogs
		WHERE title LIKE '%%%s%%'
		ORDER BY %s %s
		LIMIT %d OFFSET %d;`,
		queryParam.Title,
		queryParam.Sort,
		queryParam.SortType,
		queryParam.Limit,
		queryParam.Offset,
	)

	log.Println(query)

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	rows, err := s.db.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, err
	}

	var blogs []Blog
	for rows.Next() {
		blog := Blog{}
		rows.Scan(
			&blog.ID,
			&blog.UserID,
			&blog.CategoryID,
			&blog.Slug,
			&blog.Title,
			&blog.Content,
			&blog.Thumbnail,
			&blog.IsPublish,
			&blog.CreatedAt,
			&blog.UpdatedAt,
		)
		blogs = append(blogs, blog)
	}

	return &blogs, nil
}

func (s *BlogStore) GetBySlug(ctx context.Context, slug string) (*Blog, error) {
	query := `
	SELECT id, user_id, category_id, slug, title, content, thumbnail, is_publish, created_at, updated_at
	FROM blogs WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	rows := s.db.QueryRowContext(
		ctx,
		query,
		slug,
	)

	var blog Blog
	err := rows.Scan(
		&blog.ID,
		&blog.UserID,
		&blog.CategoryID,
		&blog.Slug,
		&blog.Title,
		&blog.Content,
		&blog.Thumbnail,
		&blog.IsPublish,
		&blog.CreatedAt,
		&blog.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &blog, nil
}

func (s *BlogStore) Block(ctx context.Context, blockedBlog BlockedBlog) error {
	query := `
		INSERT INTO blocked_blogs(id, blog_id, reasong)
		VALUES($1, $2, $3) RETURNING created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	rows := s.db.QueryRowContext(
		ctx,
		query,
		blockedBlog.ID,
		blockedBlog.BlogID,
		blockedBlog.Reason,
	)

	err := rows.Scan(
		&blockedBlog.CreatedAt,
		&blockedBlog.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
