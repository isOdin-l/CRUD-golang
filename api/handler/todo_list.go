package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/isOdin/RestApi/pkg/service"
)

type List struct {
	service service.TodoList
}

func NewListHandler(service service.TodoList) *List {
	return &List{service: service}
}

func (h *List) CreateList(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	var reqList databaseTypes.TodoList
	if err := json.NewDecoder(r.Body).Decode(&reqList); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listId, err := h.service.CreateList(userId, reqList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{
		"listId": listId,
	})

}

func (h *List) GetAllLists(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	lists, err := h.service.GetAllLists(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// render.JSON(w, r, map[string]interface{}{
	// 	"listId": listId,
	// })
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
