package service

import (
	"prtl-base-api/internal/domain/model"
	"prtl-base-api/internal/domain/repository"
)

type UserService struct {
    userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
    return &UserService{
        userRepository: userRepository,
    }
}

func (s *UserService) FindUserById(id int) (*model.User, error) {
    return s.userRepository.FindByID(id)
}

func (s *UserService) FindUserByUsername(username string) (*model.User, error) {
    return s.userRepository.FindByUsername(username)
}
