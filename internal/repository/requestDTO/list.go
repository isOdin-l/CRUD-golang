package requestDTO

import "github.com/google/uuid"

type CreateList struct {
	UserId      uuid.UUID
	Title       string
	Description string
}

type GetListById struct {
	UserId uuid.UUID
	ListId uuid.UUID
}
type DeleteList struct {
	UserId uuid.UUID
	ListId uuid.UUID
}
type UpdateList struct {
	SetArgs        *[]interface{}
	SetValuesQuery string
	ArgId          int
}
