package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/isOdin/RestApi/pkg/service"
)

type List struct {
	service service.TodoList
}

func NewListHandler(service service.TodoList) *List {
	return &List{service: service}
}

func (h *List) CreateList(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId")
	render.JSON(w, r, map[string]interface{}{
		"userId": userId,
	})
}

func (h *List) GetAllLists(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get All Lists"))
}

func (h *List) GetListById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getListById"))
}

func (h *List) UpdateList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateList"))
}

func (h *List) DeleteList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteList"))
}
