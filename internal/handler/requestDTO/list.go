package requestDTO

import (
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/service/requestDTO"
)

type CreateList struct {
	UserId      uuid.UUID `json:"user_id" form:"userId" validate:"required"`
	Title       string    `json:"title" form:"title" validate:"required"`
	Description string    `json:"description" form:"description"`
}

type GetListById struct {
	UserId uuid.UUID `json:"user_id" form:"userId" validate:"required"`
	ListId uuid.UUID `json:"list_id" form:"list_id"`
}

type DeleteList struct {
	UserId uuid.UUID `json:"user_id" form:"userId" validate:"required"`
	ListId uuid.UUID `json:"list_id" form:"list_id"`
}

type UpdateList struct {
	UserId      uuid.UUID `json:"user_id" form:"userId" validate:"required"`
	ListId      uuid.UUID `json:"list_id" form:"list_id"`
	Title       string    `json:"title" form:"title"`
	Description string    `json:"description" form:"title"`
}

func (m *CreateList) ToServiceModel() *requestDTO.CreateList {
	return &requestDTO.CreateList{
		UserId:      m.UserId,
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

func (m *UpdateList) ToServiceModel() *requestDTO.UpdateList {
	return &requestDTO.UpdateList{
		UserId:      m.UserId,
		ListId:      m.ListId,
		Title:       m.Title,
		Description: m.Description,
	}
}
