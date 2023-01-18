package user_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todo_project.com/internal/domain/user"
	"todo_project.com/internal/infra/memory"
)

func TestLogin(t *testing.T) {
	t.Run("should login a user_usecase", func(t *testing.T) {
		repo := &memory.UserRepository{}
		usecase := NewLogin(repo)

		email := "user_usecase@gmail.com"
		password := "123456"
		user, _ := user.NewUser("user_usecase", email, password)
		user, _ = repo.Insert(*user)

		input := LoginInput{
			Email:    email,
			Password: password,
		}

		output, err := usecase.Handle(input)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, user.ID, output.ID)
		assert.Equal(t, user.Email.String(), output.Email)
	})

	t.Run("should return a invalid login with email and password error", func(t *testing.T) {
		repo := &memory.UserRepository{}
		usecase := NewLogin(repo)

		email := "user_usecase@gmail.com"
		password := "123456"
		user, _ := user.NewUser("user_usecase", email, password)
		user, _ = repo.Insert(*user)

		input := LoginInput{
			Email:    email,
			Password: "other passwrod",
		}

		output, err := usecase.Handle(input)

		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, err, ErrUserEmailPasswordWrong)

		input = LoginInput{
			Email:    "wrongEmail@.com",
			Password: password,
		}

		output, err = usecase.Handle(input)
		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, err, ErrUserEmailPasswordWrong)
	})
}
