package handler

import (
	"net/http"

	"github.com/isOdin/RestApi/pkg/service"
)

type Auth struct {
	service *service.Service
}

func NewAuthHandler(service *service.Service) *Auth {
	return &Auth{service: service}
}

func (a *Auth) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sign Up"))
}

func (a *Auth) SignInHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sign In"))
}
