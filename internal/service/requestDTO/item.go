package requestDTO

import (
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/repository/requestDTO"
)

type CreateItem struct {
	UserId      uuid.UUID
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
	UserId      uuid.UUID
	ItemId      uuid.UUID
	Title       string
	Description string
	Done        *bool
}

func (m *CreateItem) ToRepoModelGetListById() *requestDTO.GetListById {
	return &requestDTO.GetListById{
		UserId: m.UserId,
		ListId: m.ListId,
	}
}

func (m *CreateItem) ToRepoModelCreateItem() *requestDTO.CreateItem {
	return &requestDTO.CreateItem{
		ListId:      m.ListId,
		Title:       m.Title,
		Description: m.Description,
	}
}

func (m *GetItemById) ToRepoModelGetItemById() *requestDTO.GetItemById {
	return &requestDTO.GetItemById{
		UserId: m.UserId,
		ItemId: m.ItemId,
	}
}

func (m *DeleteItem) ToRepoModelDeleteItem() *requestDTO.DeleteItem {
	return &requestDTO.DeleteItem{
		UserId: m.UserId,
		ItemId: m.ItemId,
	}
}

func (m *UpdateItem) ToRepoModelUpdateItem(setArgs *[]interface{}, setValuesQuery string, argId int) *requestDTO.UpdateItem {
	return &requestDTO.UpdateItem{
		SetArgs:        setArgs,
		SetValuesQuery: setValuesQuery,
		ArgId:          argId,
	}
}
