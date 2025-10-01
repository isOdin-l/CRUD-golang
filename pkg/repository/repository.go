package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{}
}
