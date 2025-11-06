package responseDTO

import (
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/service/responseDTO"
)

type GetItem struct {
	ItemId      uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Done        bool      `db:"done"`
}

type GetItemById struct {
	ItemId      uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Done        bool      `db:"done"`
}

func (m *GetItem) ToServiceModelGetItem() *responseDTO.GetItem {
	return &responseDTO.GetItem{
		ItemId:      m.ItemId,
		Title:       m.Title,
		Description: m.Description,
		Done:        m.Done,
	}
}

func (m *GetItemById) ToServiceModelGetItemById() *responseDTO.GetItemById {
	return &responseDTO.GetItemById{
		ItemId:      m.ItemId,
		Title:       m.Title,
		Description: m.Description,
		Done:        m.Done,
	}
}
