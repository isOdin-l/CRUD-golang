package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/isOdin/RestApi/internal/types/authTokenTypes"
	"github.com/spf13/viper"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Получаем токен из заголовка Authorization
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
				return
			}

			// Проверяем формат: "Bearer <token>"
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			// Парсим токен
			userId, err := parseJWTtoken(tokenString)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			// Записываем в контекст, чтобы другие хэндлеры/мидлвейры могли работать с данными
			r = r.WithContext(context.WithValue(r.Context(), "userId", userId))
			next.ServeHTTP(w, r)
		},
	)
}

func parseJWTtoken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &authTokenTypes.TokenClaims{}, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, jwt.ErrInvalidKeyType
		}

		return []byte(viper.GetString("JWT_SIGNING_KEY")), nil
	})

	if err != nil {
		return -1, err
	}
	claims, ok := token.Claims.(*authTokenTypes.TokenClaims)
	if !ok {
		return -1, jwt.ErrTokenInvalidClaims
	}

	return claims.UserId, nil
}
