package store

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	ErrNotFound         = errors.New("resource not found")
	QueryContextTimeout = 5 * time.Second
)

type Storage struct {
	Category interface {
		GetAllBlogCategory(context.Context, CategoryQueryParam) (*[]Category, error)
		GetAllProjectCategory(context.Context, CategoryQueryParam) (*[]Category, error)
		Create(context.Context, *Category) error
		GetByID(context.Context, string) (*Category, error)
		Update(context.Context, string, *Category) error
		Delete(context.Context, string) error
	}
	Blog interface {
		GetAll(context.Context, BlogQueryParam) (*[]Blog, error)
		GetBySlug(context.Context, string) (*Blog, error)
		Block(context.Context, BlockedBlog) error
	}
	Project interface{}
	User    interface{}
}

func NewStorage(db *sqlx.DB) Storage {
	return Storage{
		Category: &CategoryStore{db},
		Blog:     &BlogStore{db},
		Project:  &ProjectStore{db},
		User:     &UserStore{db},
	}
}
