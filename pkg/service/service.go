package service

import (
	"github.com/isOdin/RestApi/internal/storage/structure"
	"github.com/isOdin/RestApi/pkg/repository"
)

type Authorization interface {
	CreateUser(user structure.User) (int, error)
}

type TodoList interface {
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
	}
}
