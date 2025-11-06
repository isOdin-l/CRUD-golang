package requestDTO

import (
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/service/requestDTO"
)

type CreateList struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
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
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (m *CreateList) ToServiceModel(userId uuid.UUID) *requestDTO.CreateList {
	return &requestDTO.CreateList{
		UserId:      userId,
		Title:       m.Title,
		Description: m.Description,
	}
}

func (m *GetListById) ToServiceModel() *requestDTO.GetListById {
	return &requestDTO.GetListById{
		UserId: m.UserId,
		ListId: m.ListId,
	}
}

func (m *DeleteList) ToServiceModel() *requestDTO.DeleteList {
	return &requestDTO.DeleteList{
		UserId: m.UserId,
		ListId: m.ListId,
	}
}

func (m *UpdateList) ToServiceModel(userId, listId uuid.UUID) *requestDTO.UpdateList {
	return &requestDTO.UpdateList{
		UserId:      userId,
		ListId:      listId,
		Title:       m.Title,
		Description: m.Description,
	}
}
