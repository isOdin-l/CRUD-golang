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

type Item struct {
	service service.TodoItem
}

func NewItemHandler(service service.TodoItem) *Item {
	return &Item{service: service}
}

func (h *Item) CreateItem(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	listId, err := strconv.Atoi(chi.URLParam(r, "list_id"))
	if err != nil {
		http.Error(w, "Invalid id param", http.StatusInternalServerError)
		return
	}

	var reqItem databaseTypes.TodoItem
	if err := json.NewDecoder(r.Body).Decode(&reqItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	itemId, err := h.service.CreateItem(userId, listId, reqItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"itemId": itemId,
	})
}

func (h *Item) GetAllItems(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}
	items, err := h.service.GetAllItems(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"items": *items,
	})
}

func (h *Item) GetItemById(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	itemId, err := strconv.Atoi(chi.URLParam(r, "item_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	item, err := h.service.GetItemById(userId, itemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"item": *item,
	})
}

func (h *Item) UpdateItem(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	itemId, err := strconv.Atoi(chi.URLParam(r, "item_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updItem reqTypes.UpdateItem
	if err := json.NewDecoder(r.Body).Decode(&updItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.UpdateItem(userId, itemId, &updItem)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{})
}

func (h *Item) DeleteItem(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	itemId, err := strconv.Atoi(chi.URLParam(r, "item_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.DeleteItem(userId, itemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{})
}
