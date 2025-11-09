package handler

import (
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	*Auth
	*List
	*Item
}

type ServiceInterface interface {
	AuthServiceInterface
	ListServiceInterface
	ItemServiceInterface
}

func NewHandler(validate *validator.Validate, service ServiceInterface) *Handler {
	return &Handler{
		Auth: NewAuthHandler(validate, service),
		List: NewListHandler(validate, service),
		Item: NewItemHandler(validate, service),
	}
}
