package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user databaseTypes.User) (uuid.UUID, error) {
	var id uuid.UUID

	// TODO: change to ORM
	queryString := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", databaseTypes.TableUsers)
	row := r.db.QueryRow(context.Background(), queryString, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *AuthRepository) GetUser(username, password_hash string) (databaseTypes.User, error) {
	var user databaseTypes.User

	queryString := fmt.Sprintf("SELECT id, name, username, password_hash FROM %s WHERE username = $1 AND password_hash = $2 LIMIT 1", databaseTypes.TableUsers)
	err := r.db.QueryRow(context.Background(), queryString, username, password_hash).Scan(&user.Id, &user.Name, &user.Username, &user.Password)

	return user, err
}
