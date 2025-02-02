package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	db *sqlx.DB
}

var (
	UserRoleAdmin = "admin"
	UserRoleUser  = "user"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Password       string    `json:"-"`
	ProfilePicture string    `json:"profile_picture"`
	Bio            string    `json:"bio"`
	Role           string    `json:"role"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO(name, username, email, password, profile_picture, bio, role)
		VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at`

	err := s.db.QueryRowxContext(
		ctx,
		query,
		user.Name,
		user.Username,
		user.Email,
		user.Password,
		user.ProfilePicture,
		user.Bio,
		user.Role,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return err
}
