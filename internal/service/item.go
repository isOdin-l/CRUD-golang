package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/repository"
	"github.com/isOdin/RestApi/internal/service/requestDTO"
	"github.com/isOdin/RestApi/internal/service/responseDTO"
)

type TodoItemService struct {
	repoItem repository.TodoItem
	repoList repository.TodoList
}

func NewTodoItemService(repoItem repository.TodoItem, repoList repository.TodoList) *TodoItemService {
	return &TodoItemService{repoItem: repoItem, repoList: repoList}
}

func (s *TodoItemService) CreateItem(itemInfo *requestDTO.CreateItem) (uuid.UUID, error) {
	_, err := s.repoList.GetListById(itemInfo.ToRepoModelGetListById())
	if err != nil {
		return uuid.Nil, err
	}

	return s.repoItem.CreateItem(itemInfo.ToRepoModelCreateItem())
}

func (s *TodoItemService) GetAllItems(userId uuid.UUID) (*[]responseDTO.GetItem, error) {
	getedItem, err := s.repoItem.GetAllItems(userId)
	if err != nil {
		return nil, err
	}

	items := make([]responseDTO.GetItem, len(*getedItem))
	for i := range len(*getedItem) {
		items[i] = *(*getedItem)[i].ToServiceModelGetItem()
	}

	return &items, nil

}
func (s *TodoItemService) GetItemById(itemInfo *requestDTO.GetItemById) (*responseDTO.GetItemById, error) {
	item, err := s.repoItem.GetItemById(itemInfo.ToRepoModelGetItemById())
	if err != nil {
		return nil, err
	}
	return item.ToServiceModelGetItemById(), nil
}

func (s *TodoItemService) DeleteItem(itemInfo *requestDTO.DeleteItem) error {
	return s.repoItem.DeleteItem(itemInfo.ToRepoModelDeleteItem())
}

func (s *TodoItemService) UpdateItem(itemInfo *requestDTO.UpdateItem) error {
	setValues := make([]string, 0)
	setArgs := make([]interface{}, 0)
	argId := 1

	if itemInfo.Title == "" && itemInfo.Description == "" && itemInfo.Done == nil {
		return errors.New("Update structure has no values")
	}

	if itemInfo.Title != "" {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		setArgs = append(setArgs, itemInfo.Title)
		argId++
	}

	if itemInfo.Description != "" {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		setArgs = append(setArgs, itemInfo.Description)
		argId++
	}

	if itemInfo.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		setArgs = append(setArgs, itemInfo.Done)
		argId++
	}

	setValuesQuery := strings.Join(setValues, ", ")
	setArgs = append(setArgs, itemInfo.ItemId, itemInfo.UserId)

	return s.repoItem.UpdateItem(itemInfo.ToRepoModelUpdateItem(&setArgs, setValuesQuery, argId))
}
