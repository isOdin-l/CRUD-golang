package databaseTypes

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name" binding:"required"`
	Username string    `json:"username" db:"username" binding:"required"`
	Password string    `json:"password" db:"password_hash" binding:"required"`
}
