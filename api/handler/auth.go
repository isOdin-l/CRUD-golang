package handler

import (
	"encoding/json"
	"net/http"

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
	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		logrus.Errorf("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := h.service.CreateUser(reqUser)
	if err != nil {
		logrus.Errorf("Bad things")
	}
}

func (h *Auth) SignInHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sign In"))
}
