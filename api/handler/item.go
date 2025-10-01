package handler

import (
	"net/http"

	"github.com/isOdin/RestApi/pkg/service"
)

type Item struct {
	service *service.Service
}

func NewItemHandler(service *service.Service) *Item {
	return &Item{service: service}
}

func (i *Item) CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create list"))
}

func (i *Item) GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get All Lists"))
}

func (i *Item) GetItemById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getListById"))
}

func (i *Item) UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateList"))
}

func (i *Item) DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteList"))
}
