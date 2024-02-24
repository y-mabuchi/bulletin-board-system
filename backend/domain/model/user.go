package model

import (
	"time"
)

type User struct {
	UserID                UserID
	Name                  UserName
	Email                 EmailAddress
	Password              UserPassword
	RegistrationTimestamp time.Time
	LastLoginTimestamp    time.Time
}

func NewUser(
	name, password, email string,
) (*User, error) {
	uName, err := NewUserName(name)
	if err != nil {
		return nil, err
	}

	eAddress, err := NewEmailAddress(email)
	if err != nil {
		return nil, err
	}

	uPass, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}

	timestamp := time.Now()

	return &User{
		UserID:                NewUserID(),
		Name:                  uName,
		Email:                 eAddress,
		Password:              uPass,
		RegistrationTimestamp: timestamp,
		LastLoginTimestamp:    timestamp,
	}, nil
}

func (u *User) ChangeName(name string) (*User, error) {
	uName, err := NewUserName(name)
	if err != nil {
		return nil, err
	}

	timestamp := time.Now()

	return &User{
		UserID:                u.UserID,
		Name:                  uName,
		Email:                 u.Email,
		Password:              u.Password,
		RegistrationTimestamp: u.RegistrationTimestamp,
		LastLoginTimestamp:    timestamp,
	}, nil
}

func (u *User) ChangeEmail(email string) (*User, error) {
	eAddress, err := NewEmailAddress(email)
	if err != nil {
		return nil, err
	}

	timestamp := time.Now()

	return &User{
		UserID:                u.UserID,
		Name:                  u.Name,
		Email:                 eAddress,
		Password:              u.Password,
		RegistrationTimestamp: u.RegistrationTimestamp,
		LastLoginTimestamp:    timestamp,
	}, nil
}

func (u *User) ChangePassword(password string) (*User, error) {
	// TODO: compare old and new password

	uPass, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}

	timestamp := time.Now()

	return &User{
		UserID:                u.UserID,
		Name:                  u.Name,
		Email:                 u.Email,
		Password:              uPass,
		RegistrationTimestamp: u.RegistrationTimestamp,
		LastLoginTimestamp:    timestamp,
	}, nil
}
