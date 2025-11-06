package requestDTO

import "github.com/isOdin/RestApi/internal/service/requestDTO"

type SignUpUser struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (m *SignUpUser) ConvertToServiceModel() *requestDTO.CreateUser {
	return &requestDTO.CreateUser{
		Name:     m.Name,
		Username: m.Username,
		Password: m.Password,
	}
}

func (m *SignInUser) ConvertToServiceModel() *requestDTO.GenerateToken {
	return &requestDTO.GenerateToken{
		Username: m.Username,
		Password: m.Password,
	}
}
