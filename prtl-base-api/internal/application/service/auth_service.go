package service

import (
	"time"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"prtl-base-api/internal/domain/repository"
)

type AuthService struct {
	tokenAuth   jwtauth.JWTAuth
	userRepo    repository.UserRepository
}

func NewAuthService(tokenAuth *jwtauth.JWTAuth, userRepo repository.UserRepository) *AuthService {
	return &AuthService{
		tokenAuth:   *tokenAuth,
		userRepo:    userRepo,
	}
}

func (s *AuthService) AuthenticateUser(username string, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil || user.Password != password {
		return "", errors.New("invalid credentials")
	}

	_, tokenString, err := s.tokenAuth.Encode(map[string]interface{}{
		"sub": string(user.ID),
		"iss": "prtl-base-api",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
