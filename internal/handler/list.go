package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/handler/requestDTO"
	"github.com/isOdin/RestApi/internal/handler/responseDTO"
	"github.com/isOdin/RestApi/internal/service"
)

type List struct {
	service service.TodoList
}

func NewListHandler(service service.TodoList) *List {
	return &List{service: service}
}

func (h *List) CreateList(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	var reqList requestDTO.CreateList
	if err := render.DecodeJSON(r.Body, &reqList); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listId, err := h.service.CreateList(reqList.ToServiceModel(userId))
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
	userId, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	listId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id param", http.StatusInternalServerError)
		return
	}

	listInfo := requestDTO.GetListById{UserId: userId, ListId: listId}

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
	// Получаем Id пользователя из контекста
	userId, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	// Получаем Id листа из параметров URL-запроса
	listId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id param", http.StatusInternalServerError)
		return
	}

	// Данные из JSON засовываем в структуру
	var reqUpdList requestDTO.UpdateList
	if err := render.DecodeJSON(r.Body, &reqUpdList); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызываем функцию сервиса и переводим структуру хэндлера к структуре сервиса
	err = h.service.UpdateList(reqUpdList.ToServiceModel(userId, listId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{})
}

func (h *List) DeleteList(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "User id not found", http.StatusInternalServerError)
		return
	}

	listId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id param", http.StatusInternalServerError)
		return
	}

	listInfo := requestDTO.DeleteList{UserId: userId, ListId: listId}

	err = h.service.DeleteList(listInfo.ToServiceModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{})
}
