package service

import (
	"github.com/y-mabuchi/bulletin-board-system/backend/domain/model"
)

type PostService interface {
	Exist(postID model.PostID) bool
}
