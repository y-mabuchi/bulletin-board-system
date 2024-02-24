package model

import (
	"time"
)

type Thread struct {
	ThreadID            ThreadID
	Title               ThreadTitle
	Content             ThreadContent
	Creator             ThreadCreator
	CreationTimestamp   time.Time
	LastUpdateTimestamp time.Time
}

func NewThread(
	title, content string,
	creator ThreadCreator,
) (*Thread, error) {
	tTitle, err := NewThreadTitle(title)
	if err != nil {
		return nil, err
	}

	tContent, err := NewThreadContent(content)
	if err != nil {
		return nil, err
	}

	timestamp := time.Now()

	return &Thread{
		ThreadID:            NewThreadID(),
		Title:               tTitle,
		Content:             tContent,
		Creator:             creator,
		CreationTimestamp:   timestamp,
		LastUpdateTimestamp: timestamp,
	}, nil
}

func (t *Thread) ChangeTitle(title string) (*Thread, error) {
	tTitle, err := NewThreadTitle(title)
	if err != nil {
		return nil, err
	}

	timestamp := time.Now()

	return &Thread{
		ThreadID:            t.ThreadID,
		Title:               tTitle,
		Content:             t.Content,
		Creator:             t.Creator,
		CreationTimestamp:   t.CreationTimestamp,
		LastUpdateTimestamp: timestamp,
	}, nil
}

func (t *Thread) ChangeContent(content string) (*Thread, error) {
	tContent, err := NewThreadContent(content)
	if err != nil {
		return nil, err
	}

	timestamp := time.Now()

	return &Thread{
		ThreadID:            t.ThreadID,
		Title:               t.Title,
		Content:             tContent,
		Creator:             t.Creator,
		CreationTimestamp:   t.CreationTimestamp,
		LastUpdateTimestamp: timestamp,
	}, nil
}
