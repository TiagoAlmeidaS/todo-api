package memory

import (
	"github.com/google/uuid"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/domain/user"
)

type UserRepository struct {
	user []user.User
}

func (r *UserRepository) GetById(id string) (*user.User, error) {
	for _, user := range r.user {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, repository.ErrUserNotFound
}

func (r *UserRepository) GetByEmail(email string) (*user.User, error) {
	for _, user := range r.user {
		if user.Email.String() == email {
			return &user, nil
		}
	}

	return nil, repository.ErrUserNotFound
}

func (r *UserRepository) Insert(user user.User) (*user.User, error) {
	user.ID = uuid.New().String()

	r.user = append(r.user, user)
	return &user, nil
}
