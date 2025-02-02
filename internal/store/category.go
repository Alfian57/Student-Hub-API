package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CategoryStore struct {
	db *sqlx.DB
}

var (
	CategoryTypeBlog    = "blog"
	CategoryTypeProject = "project"
)

type Category struct {
	ID        uuid.UUID `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *CategoryStore) GetAllBlogCategory(ctx context.Context, queryParam CategoryQueryParam) (*[]Category, error) {
	return getAll(s, ctx, queryParam, CategoryTypeBlog)
}

func (s *CategoryStore) GetAllProjectCategory(ctx context.Context, queryParam CategoryQueryParam) (*[]Category, error) {
	return getAll(s, ctx, queryParam, CategoryTypeProject)
}

func getAll(s *CategoryStore, ctx context.Context, queryParam CategoryQueryParam, categoryType string) (*[]Category, error) {
	query := fmt.Sprintf(`
		SELECT id, slug, name, type, created_at, updated_at 
		FROM categories
		WHERE type = '%s' 
		AND name LIKE '%%%s%%'
		ORDER BY %s %s
		LIMIT %d OFFSET %d;`,
		categoryType,
		queryParam.Name,
		queryParam.Sort,
		queryParam.SortType,
		queryParam.Limit,
		queryParam.Offset,
	)

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	rows, err := s.db.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, err
	}

	var categories []Category
	for rows.Next() {
		category := Category{}
		rows.Scan(
			&category.ID,
			&category.Slug,
			&category.Name,
			&category.Type,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		categories = append(categories, category)
	}

	return &categories, nil
}

func (s *CategoryStore) Create(ctx context.Context, category *Category) error {
	query := `INSERT INTO categories (id, slug, name, type) VALUES($1, $2, $3, $4) RETURNING created_at, updated_at;`

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	rows := s.db.QueryRowContext(
		ctx,
		query,
		category.ID,
		category.Slug,
		category.Name,
		category.Type,
	)
	err := rows.Scan(
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryStore) GetByID(ctx context.Context, id string) (*Category, error) {
	query := `SELECT id, slug, name, type, created_at, updated_at 
				FROM categories WHERE id = $1;`

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	var category Category

	rows := s.db.QueryRowContext(ctx, query, id)
	err := rows.Scan(
		&category.ID,
		&category.Slug,
		&category.Name,
		&category.Type,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &category, nil
}

func (s *CategoryStore) Update(ctx context.Context, slug string, category *Category) error {
	query := `UPDATE categories SET slug=$1, name=$2, type=$3, updated_at=$4 WHERE slug=$5;`

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	result, err := s.db.ExecContext(
		ctx,
		query,
		category.Slug,
		category.Name,
		category.Type,
		category.UpdatedAt,
		slug,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *CategoryStore) Delete(ctx context.Context, slug string) error {
	query := `DELETE FROM categories WHERE slug=$1;`

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	result, err := s.db.ExecContext(
		ctx,
		query,
		slug,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
