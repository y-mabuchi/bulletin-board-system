package model

import "github.com/google/uuid"

type ThreadID uuid.UUID

func NewThreadID() ThreadID {
	return ThreadID(uuid.New())
}
