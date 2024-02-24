package repository

import (
	"github.com/y-mabuchi/bulletin-board-system/backend/domain/model"
)

type ThreadRepository interface {
	GetByID(id model.ThreadID) (*model.Thread, error)
	Create(thread model.Thread) (*model.Thread, error)
	Update(thread model.Thread) (*model.Thread, error)
	Delete(thread model.Thread) error
}
