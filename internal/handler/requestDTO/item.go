package requestDTO

import (
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/service/requestDTO"
)

type CreateItem struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
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
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        *bool  `json:"done"`
}

func (m *CreateItem) ToServiceModel(userId, listId uuid.UUID) *requestDTO.CreateItem {
	return &requestDTO.CreateItem{
		UserId:      userId,
		ListId:      listId,
		Title:       m.Title,
		Description: m.Description,
	}
}

func (m *GetItemById) ToServiceModel() *requestDTO.GetItemById {
	return &requestDTO.GetItemById{
		UserId: m.UserId,
		ItemId: m.ItemId,
	}
}

func (m *DeleteItem) ToServiceModel() *requestDTO.DeleteItem {
	return &requestDTO.DeleteItem{
		UserId: m.UserId,
		ItemId: m.ItemId,
	}
}

func (m *UpdateItem) ToServiceModel(userId, itemId uuid.UUID) *requestDTO.UpdateItem {
	return &requestDTO.UpdateItem{
		UserId:      userId,
		ItemId:      itemId,
		Title:       m.Title,
		Description: m.Description,
		Done:        m.Done,
	}
}
