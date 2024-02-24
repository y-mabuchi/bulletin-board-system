package service

import (
	"github.com/y-mabuchi/bulletin-board-system/backend/domain/model"
)

type UserService interface {
	Exist(userID model.UserID) bool
}
