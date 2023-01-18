package repository

import (
	"errors"
	"todo_project.com/internal/domain/user"
)

var (
	ErrUserNotFound = errors.New("user_usecase not found")
)

type IUserRepository interface {
	Insert(user user.User) (*user.User, error)
	GetById(id string) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
}
