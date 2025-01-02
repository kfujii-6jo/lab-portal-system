package persistence

import (
	"errors"
    "prtl-base-api/internal/domain/model"
    "prtl-base-api/internal/domain/repository"
)

type UserRepositoryImpl struct {}

func NewUserRepositoryImpl() repository.UserRepository {
    return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) FindByID(id int) (*model.User, error) {
	var users = []model.User{
		{ID: 1, Username: "username", Password: "password"},
	}
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
    return nil, errors.New("user not found")
}
