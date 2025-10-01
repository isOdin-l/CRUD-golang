package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/isOdin/RestApi/internal/storage/structure"
	"github.com/isOdin/RestApi/pkg/repository"
	"github.com/spf13/viper"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user structure.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	salt := viper.GetString("SALT")

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
