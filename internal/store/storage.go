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
	Users interface {
		Create(context.Context, *User) error
	}
}

func NewStorage(db *sqlx.DB) Storage {
	return Storage{
		Category: &CategoryStore{db},
		Users:    &UserStore{db},
	}
}
