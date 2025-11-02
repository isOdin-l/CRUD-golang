package databaseTypes

import "github.com/google/uuid"

type TodoList struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
}

type UserList struct {
	Id     uuid.UUID
	UserId uuid.UUID
	ListId uuid.UUID
}

type TodoItem struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
}

type ListsItem struct {
	Id     uuid.UUID
	ListId uuid.UUID
	ItemId uuid.UUID
}
