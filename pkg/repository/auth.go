package repository

import (
	"context"
	"fmt"

	"github.com/isOdin/RestApi/internal/storage/postgresql"
	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user databaseTypes.User) (int, error) {
	var id int

	// TODO: change to ORM
	queryString := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", postgresql.UsersTable)
	row := r.db.QueryRow(context.Background(), queryString, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *AuthRepository) GetUser(username, password_hash string) (databaseTypes.User, error) {
	var user databaseTypes.User

	queryString := fmt.Sprintf("SELECT id, name, username, password_hash FROM %s WHERE username = $1 AND password_hash = $2 LIMIT 1", postgresql.UsersTable)
	err := r.db.QueryRow(context.Background(), queryString, username, password_hash).Scan(&user.Id, &user.Name, &user.Username, &user.Password)

	return user, err
}
