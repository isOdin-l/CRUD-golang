package requestDTO

import "github.com/isOdin/RestApi/internal/service/requestDTO"

type SignUpUser struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type SignInUser struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
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
