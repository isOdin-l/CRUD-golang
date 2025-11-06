package responseDTO

import "github.com/google/uuid"

type GetList struct {
	Id          uuid.UUID
	Title       string
	Description string
}

type GetListById struct {
	Id          uuid.UUID
	Title       string
	Description string
}
