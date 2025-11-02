package service

import (
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/isOdin/RestApi/internal/types/reqTypes"
	"github.com/isOdin/RestApi/pkg/repository"
)

type Authorization interface {
	CreateUser(user databaseTypes.User) (uuid.UUID, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
	CreateList(userId uuid.UUID, list databaseTypes.TodoList) (uuid.UUID, error)
	GetAllLists(userId uuid.UUID) (*[]databaseTypes.TodoList, error)
	GetListById(userId, listId uuid.UUID) (*databaseTypes.TodoList, error)
	DeleteList(userId, listId uuid.UUID) error
	UpdateList(userId, listId uuid.UUID, updList reqTypes.UpdateList) error
}

type TodoItem interface {
	CreateItem(userId, listId uuid.UUID, item databaseTypes.TodoItem) (uuid.UUID, error)
	GetAllItems(userId uuid.UUID) (*[]databaseTypes.TodoItem, error)
	GetItemById(userId, itemId uuid.UUID) (*databaseTypes.TodoItem, error)
	DeleteItem(userId, itemId uuid.UUID) error
	UpdateItem(userId, itemId uuid.UUID, updItem *reqTypes.UpdateItem) error
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
