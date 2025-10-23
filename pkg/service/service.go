package service

import (
	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/isOdin/RestApi/internal/types/reqTypes"
	"github.com/isOdin/RestApi/pkg/repository"
)

type Authorization interface {
	CreateUser(user databaseTypes.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
	CreateList(userId int, list databaseTypes.TodoList) (int, error)
	GetAllLists(userId int) (*[]databaseTypes.TodoList, error)
	GetListById(userId, listId int) (*databaseTypes.TodoList, error)
	DeleteList(userId, listId int) error
	UpdateList(userId, listId int, updList reqTypes.UpdateList) error
}

type TodoItem interface {
	CreateItem(userId, listId int, item databaseTypes.TodoItem) (int, error)
	GetAllItems(userId int) (*[]databaseTypes.TodoItem, error)
	GetItemById(userId, itemId int) (*databaseTypes.TodoItem, error)
	DeleteItem(userId, itemId int) error
	UpdateItem(userId, itemId int, updItem *reqTypes.UpdateItem) error
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
