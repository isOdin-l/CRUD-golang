package responseDTO

import (
	"github.com/google/uuid"
)

type GetedUser struct {
	Id           uuid.UUID
	Name         string
	Username     string
	PasswordHash string
}
