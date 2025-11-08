package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/isOdin/RestApi/internal/handler/requestDTO"
	"github.com/isOdin/RestApi/internal/service"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	validate *validator.Validate
	service  service.Authorization
}

func NewAuthHandler(validate *validator.Validate, service service.Authorization) *Auth {
	return &Auth{validate: validate, service: service}
}

func (h *Auth) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var reqUser requestDTO.SignUpUser
	if err := render.DecodeJSON(r.Body, &reqUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(reqUser); err != nil {
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
	if err := render.DecodeJSON(r.Body, &reqUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(reqUser); err != nil {
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
