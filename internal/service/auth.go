package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/isOdin/RestApi/configs"
	jwtToken "github.com/isOdin/RestApi/internal/middleware/dto"
	"github.com/isOdin/RestApi/internal/repository"
	"github.com/isOdin/RestApi/internal/service/requestDTO"
)

type AuthService struct {
	repo repository.Authorization
	cfg  *configs.InternalConfig
}

const tokenTTL = 12 * time.Hour

func NewAuthService(cfg *configs.InternalConfig, repo repository.Authorization) *AuthService {
	return &AuthService{cfg: cfg, repo: repo}
}

func (s *AuthService) CreateUser(user *requestDTO.CreateUser) (uuid.UUID, error) {
	return s.repo.CreateUser(user.ConvertToRepoModel(s.generatePasswordHash(user.Password)))
}

func (s *AuthService) GenerateToken(user *requestDTO.GenerateToken) (string, error) {
	userFromDB, err := s.repo.GetUser(user.ConvertToRepoModel(s.generatePasswordHash(user.Password)))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtToken.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: userFromDB.Id,
	})

	return token.SignedString([]byte(s.cfg.JWT_SIGNING_KEY))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.cfg.SALT)))
}
