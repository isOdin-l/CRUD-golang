package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/internal/types/authTokenTypes"
	"github.com/isOdin/RestApi/internal/types/databaseTypes"
	"github.com/isOdin/RestApi/pkg/repository"
	"github.com/spf13/viper"
)

type AuthService struct {
	repo repository.Authorization
}

const tokenTTL = 12 * time.Hour

// TODO: move to another directory with types and structuers

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user databaseTypes.User) (uuid.UUID, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &authTokenTypes.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(viper.GetString("JWT_SIGNING_KEY")))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	salt := viper.GetString("SALT")

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
