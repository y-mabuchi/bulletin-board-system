package repository

import "github.com/y-mabuchi/bulletin-board-system/backend/domain/model"

type UserRepository interface {
	GetByID(id model.UserID) (*model.User, error)
	Create(user model.User) (*model.User, error)
	Update(user model.User) (*model.User, error)
	Delete(user model.User) error
}
