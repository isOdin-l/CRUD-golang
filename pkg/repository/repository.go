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
	GetListById(userId, listId int) (*databaseTypes.TodoList, error)
	DeleteList(userId, listId int) error
	UpdateList(setArgs *[]interface{}, argId int, setValuesQuery *string) error
}

type TodoItem interface {
	CreateItem(listId int, item databaseTypes.TodoItem) (int, error)
	GetAllItems(userId int) (*[]databaseTypes.TodoItem, error)
	GetItemById(userId, itemId int) (*databaseTypes.TodoItem, error)
	DeleteItem(userId, itemId int) error
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
