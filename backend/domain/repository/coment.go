package repository

import "github.com/y-mabuchi/bulletin-board-system/backend/domain/model"

type CommentRepository interface {
	Create(comment model.Comment) (*model.Comment, error)
	Update(comment model.Comment) (*model.Comment, error)
	Delete(comment model.Comment) error
}
