package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/handler/requestDTO"
	servReqDTO "github.com/isOdin/RestApi/internal/service/requestDTO"
	"github.com/isOdin/RestApi/tools/chiBinding"
	"github.com/sirupsen/logrus"
)

type AuthServiceInterface interface {
	CreateUser(user *servReqDTO.CreateUser) (uuid.UUID, error)
	GenerateToken(user *servReqDTO.GenerateToken) (string, error)
}

type Auth struct {
	validate *validator.Validate
	service  AuthServiceInterface
}

func NewAuthHandler(validate *validator.Validate, service AuthServiceInterface) *Auth {
	return &Auth{validate: validate, service: service}
}

func (h *Auth) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var reqUser requestDTO.SignUpUser
	if err := chiBinding.BindValidate(r, &reqUser, h.validate); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userId, err := h.service.CreateUser(reqUser.ConvertToServiceModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"id": userId,
	})
}

func (h *Auth) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var reqUser requestDTO.SignInUser
	if err := chiBinding.BindValidate(r, &reqUser, h.validate); err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	generatedToken, err := h.service.GenerateToken(reqUser.ConvertToServiceModel())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"token": generatedToken,
	})
}
