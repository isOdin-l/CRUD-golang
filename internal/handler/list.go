package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/handler/requestDTO"
	"github.com/isOdin/RestApi/internal/handler/responseDTO"
	"github.com/isOdin/RestApi/internal/service"
	"github.com/isOdin/RestApi/tools/chiBinding"
	"github.com/sirupsen/logrus"
)

type List struct {
	validate *validator.Validate
	service  service.TodoList
}

func NewListHandler(validate *validator.Validate, service service.TodoList) *List {
	return &List{validate: validate, service: service}
}

func (h *List) CreateList(w http.ResponseWriter, r *http.Request) {
	var reqList requestDTO.CreateList
	if err := chiBinding.DefaultBind(r, &reqList); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(reqList); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	listId, err := h.service.CreateList(reqList.ToServiceModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"listId": listId,
	})
}

func (h *List) GetAllLists(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	listsResponsed, err := h.service.GetAllLists(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lists := make([]responseDTO.GetList, len(*listsResponsed))
	for i := range len(*listsResponsed) {
		// ------- Указатель на массив -> массив -> элемент массива -> перевод элемента в указатель на другой тип -> элемент другого типа -------
		lists[i] = *((*listsResponsed)[i].ToHandlerModel())
	}

	render.JSON(w, r, map[string]interface{}{
		"lists": lists,
	})
}

func (h *List) GetListById(w http.ResponseWriter, r *http.Request) {
	var listInfo requestDTO.GetListById
	if err := chiBinding.DefaultBind(r, &listInfo); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.validate.Struct(listInfo); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	list, err := h.service.GetListById(listInfo.ToServiceModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{
		"list": list,
	})
}

func (h *List) UpdateList(w http.ResponseWriter, r *http.Request) {
	var reqUpdList requestDTO.UpdateList
	if err := chiBinding.DefaultBind(r, &reqUpdList); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(reqUpdList); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.UpdateList(reqUpdList.ToServiceModel()); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{})
}

func (h *List) DeleteList(w http.ResponseWriter, r *http.Request) {
	var listInfo requestDTO.DeleteList
	if err := chiBinding.DefaultBind(r, &listInfo); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.validate.Struct(listInfo); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.DeleteList(listInfo.ToServiceModel()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{})
}
