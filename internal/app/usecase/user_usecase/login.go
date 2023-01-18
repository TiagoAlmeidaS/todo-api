package user_usecase

import (
	"errors"
	"todo_project.com/internal/app/repository"
)

var (
	ErrUserEmailPasswordWrong = errors.New("email or password is wrong")
)

type Login interface {
	Handle(input LoginInput) (*Output, error)
}

type login struct {
	userRepository repository.IUserRepository
}

func NewLogin(userRepository repository.IUserRepository) Login {
	return &login{userRepository: userRepository}
}

type LoginInput struct {
	Email    string
	Password string
}

func (u *login) Handle(input LoginInput) (*Output, error) {
	user, err := u.userRepository.GetByEmail(input.Email)

	if err != nil {
		if err == repository.ErrUserNotFound {
			return nil, ErrUserEmailPasswordWrong
		}
		return nil, err
	}

	if !user.Password.IsCorrectPassword(input.Password) {
		return nil, ErrUserEmailPasswordWrong
	}

	return userOutputFromUser(user), nil
}
