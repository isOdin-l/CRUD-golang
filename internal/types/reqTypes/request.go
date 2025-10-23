package reqTypes

type UpdateList struct {
	Title       *string
	Description *string
}

type UpdateItem struct {
	Title       *string
	Description *string
	Done        *bool
}
