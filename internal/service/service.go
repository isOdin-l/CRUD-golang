package service

import (
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/configs"
	"github.com/isOdin/RestApi/internal/repository"
	"github.com/isOdin/RestApi/internal/service/requestDTO"
	"github.com/isOdin/RestApi/internal/service/responseDTO"
)

type Authorization interface {
	CreateUser(user *requestDTO.CreateUser) (uuid.UUID, error)
	GenerateToken(user *requestDTO.GenerateToken) (string, error)
}

type TodoList interface {
	CreateList(listInfo *requestDTO.CreateList) (uuid.UUID, error)
	GetAllLists(userId uuid.UUID) (*[]responseDTO.GetList, error)
	GetListById(listInfo *requestDTO.GetListById) (*responseDTO.GetListById, error)
	DeleteList(listInfo *requestDTO.DeleteList) error
	UpdateList(listInfo *requestDTO.UpdateList) error
}

type TodoItem interface {
	CreateItem(itemInfo *requestDTO.CreateItem) (uuid.UUID, error)
	GetAllItems(userId uuid.UUID) (*[]responseDTO.GetItem, error)
	GetItemById(itemInfo *requestDTO.GetItemById) (*responseDTO.GetItemById, error)
	DeleteItem(itemInfo *requestDTO.DeleteItem) error
	UpdateItem(itemInfo *requestDTO.UpdateItem) error
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(cfg *configs.InternalConfig, repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(cfg, repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
