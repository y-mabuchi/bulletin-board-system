package model

import "github.com/google/uuid"

type PostID uuid.UUID

func NewPostID() PostID {
	return PostID(uuid.New())
}
