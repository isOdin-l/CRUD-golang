package repository

import (
	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Authorization interface {
	CreateUser(user databaseTypes.User) (int, error)
	GetUser(username, password string) (databaseTypes.User, error)
}

type TodoList interface {
	CreateList(userId int, list databaseTypes.TodoList) (int, error)
	GetAllLists(userId int) (*[]databaseTypes.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		TodoList:      NewTodoListRepository(db),
	}
}
