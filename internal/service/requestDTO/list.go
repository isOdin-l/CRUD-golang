package requestDTO

import (
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/repository/requestDTO"
)

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
	UserId      uuid.UUID
	ListId      uuid.UUID
	Title       string
	Description string
}

func (m *CreateList) ConvertToRepoModel() *requestDTO.CreateList {
	return &requestDTO.CreateList{
		UserId:      m.UserId,
		Title:       m.Title,
		Description: m.Description,
	}
}

func (m *GetListById) ConvertToRepoModel() *requestDTO.GetListById {
	return &requestDTO.GetListById{
		UserId: m.UserId,
		ListId: m.ListId,
	}
}

func (m *DeleteList) ConvertToRepoModel() *requestDTO.DeleteList {
	return &requestDTO.DeleteList{
		UserId: m.UserId,
		ListId: m.ListId,
	}
}

func (m *UpdateList) ConvertToRepoModel(setArgs *[]interface{}, argId int, setValuesQuery string) *requestDTO.UpdateList {
	return &requestDTO.UpdateList{
		SetValuesQuery: setValuesQuery,
		SetArgs:        setArgs,
		ArgId:          argId,
	}
}
