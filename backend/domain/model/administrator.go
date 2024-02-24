package model

import (
	"time"
)

type Administrator struct {
	AdministratorID       UserID
	Name                  UserName
	Email                 EmailAddress
	Password              UserPassword
	RegistrationTimestamp time.Time
	LastLoginTimestamp    time.Time
}

func NewAdministrator(
	name, password, email string,
) (*Administrator, error) {
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

	return &Administrator{
		AdministratorID:       NewUserID(),
		Name:                  uName,
		Email:                 eAddress,
		Password:              uPass,
		RegistrationTimestamp: timestamp,
		LastLoginTimestamp:    timestamp,
	}, nil
}

func (a *Administrator) ChangeName(name string) (*Administrator, error) {
	uName, err := NewUserName(name)
	if err != nil {
		return nil, err
	}

	return &Administrator{
		AdministratorID:       a.AdministratorID,
		Name:                  uName,
		Email:                 a.Email,
		Password:              a.Password,
		RegistrationTimestamp: a.RegistrationTimestamp,
		LastLoginTimestamp:    a.LastLoginTimestamp,
	}, nil
}

func (a *Administrator) ChangeEmail(email string) (*Administrator, error) {
	eAddress, err := NewEmailAddress(email)
	if err != nil {
		return nil, err
	}

	return &Administrator{
		AdministratorID:       a.AdministratorID,
		Name:                  a.Name,
		Email:                 eAddress,
		Password:              a.Password,
		RegistrationTimestamp: a.RegistrationTimestamp,
		LastLoginTimestamp:    time.Time{},
	}, nil
}

func (a *Administrator) ChangePassword(password string) (*Administrator, error) {
	uPass, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}

	// TODO: compare old and new password

	return &Administrator{
		AdministratorID:       a.AdministratorID,
		Name:                  a.Name,
		Email:                 a.Email,
		Password:              uPass,
		RegistrationTimestamp: a.RegistrationTimestamp,
		LastLoginTimestamp:    a.LastLoginTimestamp,
	}, nil
}
