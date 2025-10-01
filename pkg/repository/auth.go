package repository

import (
	"github.com/isOdin/RestApi/internal/storage/structure"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user structure.User) (int, error) {
	return 0, nil
}
