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

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) CreateList(userId uuid.UUID, list databaseTypes.TodoList) (uuid.UUID, error) {
	return s.repo.CreateList(userId, list)
}

func (s *TodoListService) GetAllLists(userId uuid.UUID) (*[]databaseTypes.TodoList, error) {
	return s.repo.GetAllLists(userId)
}

func (s *TodoListService) GetListById(userId, listId uuid.UUID) (*databaseTypes.TodoList, error) {
	return s.repo.GetListById(userId, listId)
}

func (s *TodoListService) DeleteList(userId, listId uuid.UUID) error {
	return s.repo.DeleteList(userId, listId)
}

func (s *TodoListService) UpdateList(userId, listId uuid.UUID, updList reqTypes.UpdateList) error {
	setValues := make([]string, 0)
	setArgs := make([]interface{}, 0)
	argId := 1

	if updList.Title == nil && updList.Description == nil {
		return errors.New("Update structure has no values")
	}

	if updList.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		setArgs = append(setArgs, *updList.Title)
		argId++
	}

	if updList.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		setArgs = append(setArgs, *updList.Description)
		argId++
	}

	setValuesQuery := strings.Join(setValues, ", ")
	setArgs = append(setArgs, listId, userId)

	return s.repo.UpdateList(&setArgs, argId, &setValuesQuery)
}
