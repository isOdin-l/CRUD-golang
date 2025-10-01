package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/isOdin/RestApi/internal/storage/structure"
	"github.com/isOdin/RestApi/pkg/service"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	service service.Authorization
}

func NewAuthHandler(service service.Authorization) *Auth {
	return &Auth{service: service}
}

func (h *Auth) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var reqUser structure.User
	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil { // TODO: переделать с render
		logrus.Errorf("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateUser(reqUser)
	if err != nil {
		logrus.Errorf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError) // TODO: cубрать дублирование кода
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"id": id,
	})
}

func (h *Auth) SignInHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sign In"))
}
