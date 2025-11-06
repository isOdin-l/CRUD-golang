package requestDTO

type CreateUser struct {
	Name         string
	Username     string
	PasswordHash string
}

type GetUser struct {
	Username     string
	PasswordHash string
}
