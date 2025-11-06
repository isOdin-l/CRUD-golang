package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/handler/requestDTO"
	"github.com/isOdin/RestApi/internal/service"
)

type Item struct {
	service service.TodoItem
}

func NewItemHandler(service service.TodoItem) *Item {
	return &Item{service: service}
}

func (h *Item) CreateItem(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	listId, err := uuid.Parse(chi.URLParam(r, "list_id"))
	if err != nil {
		http.Error(w, "Invalid id param", http.StatusInternalServerError)
		return
	}

	var reqItem requestDTO.CreateItem
	if err := json.NewDecoder(r.Body).Decode(&reqItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	itemId, err := h.service.CreateItem(reqItem.ToServiceModel(userId, listId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"itemId": itemId,
	})
}

func (h *Item) GetAllItems(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(uuid.UUID)
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
	userId, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	itemId, err := uuid.Parse(chi.URLParam(r, "item_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	itemInfo := &requestDTO.GetItemById{UserId: userId, ItemId: itemId}

	item, err := h.service.GetItemById(itemInfo.ToServiceModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"item": *item,
	})
}

func (h *Item) UpdateItem(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	itemId, err := uuid.Parse(chi.URLParam(r, "item_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updItem requestDTO.UpdateItem
	if err := json.NewDecoder(r.Body).Decode(&updItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.UpdateItem(updItem.ToServiceModel(userId, itemId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{})
}

func (h *Item) DeleteItem(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	itemId, err := uuid.Parse(chi.URLParam(r, "item_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	itemInfo := requestDTO.DeleteItem{UserId: userId, ItemId: itemId}

	err = h.service.DeleteItem(itemInfo.ToServiceModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{})
}
