package responseDTO

import "github.com/google/uuid"

type GetItem struct {
	ItemId      uuid.UUID
	Title       string
	Description string
	Done        bool
}

type GetItemById struct {
	ItemId      uuid.UUID
	Title       string
	Description string
	Done        bool
}
