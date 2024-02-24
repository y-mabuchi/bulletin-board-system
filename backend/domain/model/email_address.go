package model

import (
	"fmt"
	"net/mail"
)

type EmailAddress string

func NewEmailAddress(email string) (EmailAddress, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", fmt.Errorf(
			"failed to parse email: %s, err: %v",
			email,
			err,
		)
	}

	return EmailAddress(email), nil
}
