package requestDTO

import "github.com/isOdin/RestApi/internal/repository/requestDTO"

type CreateUser struct {
	Name     string
	Username string
	Password string
}

type GenerateToken struct {
	Username string
	Password string
}

func (m *CreateUser) ConvertToRepoModel(passwordHash string) *requestDTO.CreateUser {
	return &requestDTO.CreateUser{
		Name:         m.Name,
		Username:     m.Username,
		PasswordHash: passwordHash,
	}
}

func (m *GenerateToken) ConvertToRepoModel(passwordHash string) *requestDTO.GetUser {
	return &requestDTO.GetUser{
		Username:     m.Username,
		PasswordHash: passwordHash,
	}
}
