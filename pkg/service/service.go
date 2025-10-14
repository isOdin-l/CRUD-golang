package service

import (
	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/isOdin/RestApi/pkg/repository"
)

type Authorization interface {
	CreateUser(user databaseTypes.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
	CreateList(userId int, list databaseTypes.TodoList) (int, error)
	GetAllLists(userId int) (*[]databaseTypes.TodoList, error)
}

type TodoItem interface {
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
	}
}
