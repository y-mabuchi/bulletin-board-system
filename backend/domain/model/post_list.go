package model

import "fmt"

type PostList struct {
	ThreadID ThreadID
	Posts    []Post
}

func (pl *PostList) AddPost(post Post) (*PostList, error) {
	if ok := pl.canAddPost(); !ok {
		return nil, fmt.Errorf("posts are full")
	}

	posts := pl.Posts
	posts = append(posts, post)

	return &PostList{
		ThreadID: pl.ThreadID,
		Posts:    posts,
	}, nil
}

func (pl *PostList) canAddPost() bool {
	count := len(pl.Posts)

	if count > 999 {
		return false
	}

	return true
}
