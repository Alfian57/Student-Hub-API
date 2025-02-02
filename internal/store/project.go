package store

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProjectStore struct {
	db *sqlx.DB
}

type Project struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	CategoryID  uuid.UUID `json:"category_id"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"`
	IsPublish   bool      `json:"is_publish"`
	CodeLink    string    `json:"code_link"`
	AppLink     string    `json:"app_link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BlockedProject struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *ProjectStore) Create(ctx context.Context, project *Project) error {
	query := `
		INSERT INTO blogs (id, user_id, category_id, slug, title, description, thumbnail, is_publish, code_link, app_link) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING created_at, updated_at;`

	ctx, cancel := context.WithTimeout(ctx, QueryContextTimeout)
	defer cancel()

	rows := s.db.QueryRowContext(
		ctx,
		query,
		project.ID,
		project.UserID,
		project.CategoryID,
		project.Slug,
		project.Title,
		project.Description,
		project.Thumbnail,
		project.IsPublish,
		project.CodeLink,
		project.AppLink,
	)
	err := rows.Scan(
		&project.CreatedAt,
		&project.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
