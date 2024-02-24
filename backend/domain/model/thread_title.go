package model

import "fmt"

type ThreadTitle string

func NewThreadTitle(title string) (ThreadTitle, error) {
	if len(title) > 255 {
		return "", fmt.Errorf(
			"title must be less than or equal 255: %s",
			title,
		)
	}

	return ThreadTitle(title), nil
}
