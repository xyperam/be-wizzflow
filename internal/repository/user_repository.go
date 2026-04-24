package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xyperam/wizzflow/internal/models"
)

type UserRepository interface {
	SaveUser(ctx context.Context, u models.User) (models.User, error)
	FindUserByUsername(ctx context.Context, username string) (models.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

// save user saat register

func (r *userRepository) SaveUser(ctx context.Context, u models.User) (models.User, error) {
	query := `INSERT INTO users (username,email,password)
	VALUES ($1,$2,$3)
	RETURNING id,created_at,updated_at`

	err := r.db.QueryRow(ctx, query, u.Username, u.Email, u.Password).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return models.User{}, err
	}
	return u, nil

}
func (r *userRepository) FindUserByUsername(ctx context.Context, username string) (models.User, error) {
	var u models.User
	query := `SELECT id,username,email,password,created_at,updated_at FROM users WHERE username =%1`
	err := r.db.QueryRow(ctx, query, username).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return models.User{}, err
	}
	return u, nil
}
