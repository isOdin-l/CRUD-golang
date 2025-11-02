package repository

import (
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Authorization interface {
	CreateUser(user databaseTypes.User) (uuid.UUID, error)
	GetUser(username, password string) (databaseTypes.User, error)
}

type TodoList interface {
	CreateList(userId uuid.UUID, list databaseTypes.TodoList) (uuid.UUID, error)
	GetAllLists(userId uuid.UUID) (*[]databaseTypes.TodoList, error)
	GetListById(userId, listId uuid.UUID) (*databaseTypes.TodoList, error)
	DeleteList(userId, listId uuid.UUID) error
	UpdateList(setArgs *[]interface{}, argId int, setValuesQuery *string) error
}

type TodoItem interface {
	CreateItem(listId uuid.UUID, item databaseTypes.TodoItem) (uuid.UUID, error)
	GetAllItems(userId uuid.UUID) (*[]databaseTypes.TodoItem, error)
	GetItemById(userId, itemId uuid.UUID) (*databaseTypes.TodoItem, error)
	DeleteItem(userId, itemId uuid.UUID) error
	UpdateItem(setArgs *[]interface{}, setValuesQuery *string, argId int) error
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
		TodoItem:      NewTodoItemRepository(db),
	}
}
