package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/handler/requestDTO"
	"github.com/isOdin/RestApi/internal/service"
	"github.com/isOdin/RestApi/tools/chiBinding"
	"github.com/sirupsen/logrus"
)

type Item struct {
	validate *validator.Validate
	service  service.TodoItem
}

func NewItemHandler(validate *validator.Validate, service service.TodoItem) *Item {
	return &Item{validate: validate, service: service}
}

func (h *Item) CreateItem(w http.ResponseWriter, r *http.Request) {
	var reqItem requestDTO.CreateItem
	if err := chiBinding.DefaultBind(r.Clone(r.Context()), &reqItem); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	itemId, err := h.service.CreateItem(reqItem.ToServiceModel())
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
	var itemInfo requestDTO.GetItemById
	if err := chiBinding.DefaultBind(r, &itemInfo); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
	var updItem requestDTO.UpdateItem
	if err := chiBinding.DefaultBind(r, &updItem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.UpdateItem(updItem.ToServiceModel()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{})
}

func (h *Item) DeleteItem(w http.ResponseWriter, r *http.Request) {
	var itemInfo requestDTO.DeleteItem
	if err := chiBinding.DefaultBind(r, &itemInfo); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.DeleteItem(itemInfo.ToServiceModel()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{})
}
