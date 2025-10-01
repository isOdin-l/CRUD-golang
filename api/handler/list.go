package handler

import (
	"net/http"

	"github.com/isOdin/RestApi/pkg/service"
)

type List struct {
	service *service.Service
}

func NewListHandler(service *service.Service) *List {
	return &List{service: service}
}

func (l *List) CreateList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create list"))
}

func (l *List) GetAllLists(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get All Lists"))
}

func (l *List) GetListById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getListById"))
}

func (l *List) UpdateList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateList"))
}

func (l *List) DeleteList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteList"))
}
