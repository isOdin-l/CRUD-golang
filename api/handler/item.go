package handler

import (
	"net/http"

	"github.com/isOdin/RestApi/pkg/service"
)

type Item struct {
	service service.TodoItem
}

func NewItemHandler(service service.TodoItem) *Item {
	return &Item{service: service}
}

func (h *Item) CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create list"))
}

func (h *Item) GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get All Lists"))
}

func (h *Item) GetItemById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getListById"))
}

func (h *Item) UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateList"))
}

func (h *Item) DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteList"))
}
