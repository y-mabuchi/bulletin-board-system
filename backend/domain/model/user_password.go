package model

import "fmt"

type UserPassword string

func NewUserPassword(password string) (UserPassword, error) {
	passLen := len(password)
	if passLen < 8 || passLen > 16 {
		return "", fmt.Errorf("password must be between 8 and 16")
	}

	return UserPassword(password), nil
}
