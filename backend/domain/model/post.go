package model

import (
	"time"
)

type Post struct {
	PostID              PostID
	ThreadID            ThreadID
	Content             PostContent
	Poster              Poster
	QuotedPostID        *PostID
	PostTimestamp       time.Time
	LastUpdateTimestamp time.Time
}

func NewPost(
	threadID ThreadID,
	content string,
	poster Poster,
	quotedPostID *PostID,
) (*Post, error) {
	pContent, err := NewPostContent(content)
	if err != nil {
		return nil, err
	}

	// TODO: exist checks
	// - thread
	// - post(if quoted post id is not nil)

	timestamp := time.Now()

	return &Post{
		PostID:              NewPostID(),
		ThreadID:            threadID,
		Content:             pContent,
		Poster:              poster,
		QuotedPostID:        quotedPostID,
		PostTimestamp:       timestamp,
		LastUpdateTimestamp: timestamp,
	}, nil
}

func (p *Post) ChangeContent(content string) (*Post, error) {
	pContent, err := NewPostContent(content)
	if err != nil {
		return nil, err
	}
	timestamp := time.Now()

	return &Post{
		PostID:              p.PostID,
		ThreadID:            p.ThreadID,
		Content:             pContent,
		Poster:              p.Poster,
		QuotedPostID:        p.QuotedPostID,
		PostTimestamp:       p.PostTimestamp,
		LastUpdateTimestamp: timestamp,
	}, nil
}
