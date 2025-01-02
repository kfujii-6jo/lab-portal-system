package service

import (
	"errors"
	"prtl-base-api/internal/domain/repository"
	"prtl-base-api/internal/infrastructure/jwt"
)

type AuthService struct {
	jwtProvider jwt.JWTProvider
	userRepo    repository.UserRepository
}

func NewAuthService(jwtProvider jwt.JWTProvider, userRepo repository.UserRepository) *AuthService {
	return &AuthService{
		jwtProvider: jwtProvider,
		userRepo:    userRepo,
	}
}

func (s *AuthService) AuthenticateUser(username string, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil || user.Password != password {
		return "", errors.New("invalid credentials")
	}
	token, err := s.jwtProvider.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
