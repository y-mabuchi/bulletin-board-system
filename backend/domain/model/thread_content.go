package model

import "fmt"

type ThreadContent string

func NewThreadContent(content string) (ThreadContent, error) {
	if len(content) > 1000 {
		return "", fmt.Errorf(
			"content must be less than or equal 1000: %s",
			content,
		)
	}

	return ThreadContent(content), nil
}
