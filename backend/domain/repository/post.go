package repository

import "github.com/y-mabuchi/bulletin-board-system/backend/domain/model"

type PostRepository interface {
	GetAll(threadID model.ThreadID) ([]model.Post, error)
	Create(post model.Post) (*model.Post, error)
	Update(post model.Post) (*model.Post, error)
	Delete(post model.Post) error
}
