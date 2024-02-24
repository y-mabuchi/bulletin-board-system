package model

import "fmt"

type PostContent string

func NewPostContent(content string) (PostContent, error) {
	if len(content) > 1000 {
		return "", fmt.Errorf(
			"comment must be less than or equal 1000: %s",
			content,
		)
	}

	return PostContent(content), nil
}
