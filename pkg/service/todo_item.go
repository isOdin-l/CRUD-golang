package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/isOdin/RestApi/internal/types/reqTypes"
	"github.com/isOdin/RestApi/pkg/repository"
)

type TodoItemService struct {
	repoItem repository.TodoItem
	repoList repository.TodoList
}

func NewTodoItemService(repoItem repository.TodoItem, repoList repository.TodoList) *TodoItemService {
	return &TodoItemService{repoItem: repoItem, repoList: repoList}
}

func (s *TodoItemService) CreateItem(userId, listId uuid.UUID, item databaseTypes.TodoItem) (uuid.UUID, error) {
	_, err := s.repoList.GetListById(userId, listId)
	if err != nil {
		return uuid.Nil, err
	}

	return s.repoItem.CreateItem(listId, item)
}

func (s *TodoItemService) GetAllItems(userId uuid.UUID) (*[]databaseTypes.TodoItem, error) {
	return s.repoItem.GetAllItems(userId)
}
func (s *TodoItemService) GetItemById(userId, itemId uuid.UUID) (*databaseTypes.TodoItem, error) {
	return s.repoItem.GetItemById(userId, itemId)
}

func (s *TodoItemService) DeleteItem(userId, itemId uuid.UUID) error {
	return s.repoItem.DeleteItem(userId, itemId)
}

func (s *TodoItemService) UpdateItem(userId, itemId uuid.UUID, updItem *reqTypes.UpdateItem) error {
	setValues := make([]string, 0)
	setArgs := make([]interface{}, 0)
	argId := 1

	if updItem.Title == nil && updItem.Description == nil && updItem.Done == nil {
		return errors.New("Update structure has no values")
	}

	if updItem.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		setArgs = append(setArgs, *updItem.Title)
		argId++
	}

	if updItem.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		setArgs = append(setArgs, *updItem.Description)
		argId++
	}

	if updItem.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		setArgs = append(setArgs, *updItem.Done)
		argId++
	}

	setValuesQuery := strings.Join(setValues, ", ")
	setArgs = append(setArgs, itemId, userId)

	return s.repoItem.UpdateItem(&setArgs, &setValuesQuery, argId)
}
