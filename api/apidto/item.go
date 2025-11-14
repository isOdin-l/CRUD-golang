package apidto

type CreateItem struct {
	Title       string
	Description string
}

type UpdateItem struct {
	Title       string
	Description string
	Done        bool `validate:"optional"`
}
