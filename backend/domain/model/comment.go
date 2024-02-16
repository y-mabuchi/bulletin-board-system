package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID                   uuid.UUID
	Content              string
	PostedAt             time.Time
	Commenter            User
	RelatedBulletinBoard BulletinBoard
}

func NewComment(
	content string,
	commenter User,
	relatedBulletinBoard BulletinBoard,
) (*Comment, error) {
	if len(content) > 1000 {
		return nil, fmt.Errorf("comment must be less than or equal 1000: %s", content)
	}

	// TODO: exist check, commenter and bulletin board

	timestamp := time.Now()

	return &Comment{
		ID:                   uuid.New(),
		Content:              content,
		PostedAt:             timestamp,
		Commenter:            commenter,
		RelatedBulletinBoard: relatedBulletinBoard,
	}, nil
}
