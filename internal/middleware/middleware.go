package middleware

import (
	"net/http"

	"github.com/isOdin/RestApi/configs"
)

type Auth interface {
	JWTAuth(next http.Handler) http.Handler
}

type Middleware struct {
	Auth
}

func NewMiddleware(cfg *configs.InternalConfig) *Middleware {
	return &Middleware{
		Auth: NewAuthMiddleware(cfg),
	}
}
