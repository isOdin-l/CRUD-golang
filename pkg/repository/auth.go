package repository

import (
	"context"
	"fmt"

	"github.com/isOdin/RestApi/internal/storage/postgresql"
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
	var id int

	queryString := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", postgresql.UsersTable)
	row := r.db.QueryRow(context.Background(), queryString, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}
