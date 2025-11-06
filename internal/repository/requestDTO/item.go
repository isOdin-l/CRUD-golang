package requestDTO

import "github.com/google/uuid"

type CreateItem struct {
	ListId      uuid.UUID
	Title       string
	Description string
}

type GetItemById struct {
	UserId uuid.UUID
	ItemId uuid.UUID
}

type DeleteItem struct {
	UserId uuid.UUID
	ItemId uuid.UUID
}

type UpdateItem struct {
	SetArgs        *[]interface{}
	SetValuesQuery string
	ArgId          int
}
