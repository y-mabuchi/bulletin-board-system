package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BulletinBoard struct {
	ID        uuid.UUID
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Poster    User
}

func NewBulletinBoard(
	title, content string,
	user User,
) (*BulletinBoard, error) {
	if len(title) > 255 {
		return nil, fmt.Errorf("title must be less than or equal 255: %s", title)
	}

	if len(content) > 1000 {
		return nil, fmt.Errorf("content must be less than or equal 1000: %s", content)
	}

	// TODO: exist check, poster

	timestamp := time.Now()

	return &BulletinBoard{
		ID:        uuid.New(),
		Title:     title,
		Content:   content,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
		Poster:    user,
	}, nil
}
