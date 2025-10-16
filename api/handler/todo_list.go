package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/isOdin/RestApi/internal/types/reqTypes"
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
	render.JSON(w, r, map[string]interface{}{
		"lists": lists,
	})
}

func (h *List) GetListById(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	listId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id param", http.StatusInternalServerError)
		return
	}

	list, err := h.service.GetListById(userId, listId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{
		"list": list,
	})
}

func (h *List) UpdateList(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	listId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id param", http.StatusInternalServerError)
		return
	}

	var reqUpdList reqTypes.UpdateList
	if err := json.NewDecoder(r.Body).Decode(&reqUpdList); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateList(userId, listId, reqUpdList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{})
}

func (h *List) DeleteList(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	listId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id param", http.StatusInternalServerError)
		return
	}

	err = h.service.DeleteList(userId, listId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{})
}
