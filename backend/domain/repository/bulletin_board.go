package repository

import "github.com/y-mabuchi/bulletin-board-system/backend/domain/model"

type BulletinBoardRepository interface {
	Create(bulletinBoard model.BulletinBoard) (*model.BulletinBoard, error)
	Update(bulletinBoard model.BulletinBoard) (*model.BulletinBoard, error)
	Delete(bulletinBoard model.BulletinBoard) error
}
