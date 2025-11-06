package responseDTO

import (
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/handler/responseDTO"
)

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

func (m *GetList) ToHandlerModel() *responseDTO.GetList {
	return &responseDTO.GetList{
		Id:          m.Id,
		Title:       m.Title,
		Description: m.Description,
	}
}

func (m *GetListById) ToHandlerModel() *responseDTO.GetListById {
	return &responseDTO.GetListById{
		Id:          m.Id,
		Title:       m.Title,
		Description: m.Description,
	}
}
