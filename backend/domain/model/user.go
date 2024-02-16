package model

import (
	"fmt"
	"net/mail"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(
	name, password, email string,
) (*User, error) {
	if len(name) > 255 {
		return nil, fmt.Errorf("name must be less than or equal 255: %s", name)
	}

	passLen := len(password)
	if passLen < 8 || passLen > 16 {
		return nil, fmt.Errorf("password must be between 8 and 16")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return nil, fmt.Errorf("failed to parse email: %s, %v", email, err)
	}

	timestamp := time.Now()

	return &User{
		ID:        uuid.New(),
		Name:      name,
		Password:  password,
		Email:     email,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}, nil
}
