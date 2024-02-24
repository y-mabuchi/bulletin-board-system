package model

import (
	"fmt"
	"time"
)

const MaxPost = 999

type Thread struct {
	ThreadID            ThreadID
	Title               ThreadTitle
	Content             ThreadContent
	Posts               []Post
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

	var posts []Post

	timestamp := time.Now()

	return &Thread{
		ThreadID:            NewThreadID(),
		Title:               tTitle,
		Content:             tContent,
		Creator:             creator,
		Posts:               posts,
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

func (t *Thread) AddPost(p Post) (*Thread, error) {
	count := len(t.Posts)
	if count > MaxPost {
		return nil, fmt.Errorf("posts are full, count: %d", count)
	}

	posts := t.Posts
	posts = append(posts, p)

	return &Thread{
		ThreadID:            t.ThreadID,
		Title:               t.Title,
		Content:             t.Content,
		Posts:               posts,
		Creator:             t.Creator,
		CreationTimestamp:   t.CreationTimestamp,
		LastUpdateTimestamp: t.LastUpdateTimestamp,
	}, nil
}
