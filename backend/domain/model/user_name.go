package model

import "fmt"

type UserName string

func NewUserName(name string) (UserName, error) {
	if len(name) > 255 {
		return "", fmt.Errorf(
			"name must be less than or equal 255: %s",
			name,
		)
	}
	return UserName(name), nil
}
