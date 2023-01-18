package user

import (
	"errors"
)

var (
	ErrNameIsInvalid = errors.New("name don't to be empty")
)

type User struct {
	ID       string
	Name     string
	Password Password
	Email    Email
}

func NewUser(name, email, password string) (*User, error) {
	if name == "" {
		return nil, ErrNameIsInvalid
	}

	emailRes, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	passwordRes, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		Email:    *emailRes,
		Password: *passwordRes,
	}, nil
}
