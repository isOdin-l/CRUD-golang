package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/isOdin/RestApi/internal/service"
)

type Authorization interface {
	SignUpHandler(w http.ResponseWriter, r *http.Request)
	SignInHandler(w http.ResponseWriter, r *http.Request)
}

type TodoItem interface {
	CreateItem(w http.ResponseWriter, r *http.Request)
	GetAllItems(w http.ResponseWriter, r *http.Request)
	GetItemById(w http.ResponseWriter, r *http.Request)
	UpdateItem(w http.ResponseWriter, r *http.Request)
	DeleteItem(w http.ResponseWriter, r *http.Request)
}

type TodoList interface {
	CreateList(w http.ResponseWriter, r *http.Request)
	GetAllLists(w http.ResponseWriter, r *http.Request)
	GetListById(w http.ResponseWriter, r *http.Request)
	UpdateList(w http.ResponseWriter, r *http.Request)
	DeleteList(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	Authorization
	TodoItem
	TodoList
}

func NewHandler(validate *validator.Validate, service *service.Service) *Handler {
	return &Handler{
		Authorization: NewAuthHandler(validate, service.Authorization),
		TodoItem:      NewItemHandler(validate, service.TodoItem),
		TodoList:      NewListHandler(validate, service.TodoList),
	}
}
