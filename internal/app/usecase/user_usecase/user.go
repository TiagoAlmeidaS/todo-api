package user_usecase

import (
	"todo_project.com/internal/domain/user"
)

type Output struct {
	ID    string
	Name  string
	Email string
}

func userOutputFromUser(user *user.User) *Output {
	return &Output{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email.String(),
	}
}
