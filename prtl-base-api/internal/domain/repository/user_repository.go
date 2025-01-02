package repository

import (
	"prtl-base-api/internal/domain/model"
)

type UserRepository interface {
    FindByID(id int) (*model.User, error)
}
