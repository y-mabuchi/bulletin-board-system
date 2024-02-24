package model

import "github.com/google/uuid"

type UserID uuid.UUID

func NewUserID() UserID {
	uid := uuid.New()

	return UserID(uid)
}
