package store

import (
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
	IsPublish   string    `json:"is_publish"`
	CodeLink    string    `json:"code_link"`
	AppLink     string    `json:"app_link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
