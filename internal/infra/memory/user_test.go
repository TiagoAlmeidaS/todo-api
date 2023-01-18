package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/domain/user"
)

func TestUserRepository_GetById(t *testing.T) {

	userPassword := "1234556"
	userFix, _ := user.NewUser("tiago", "tiagotigore@hotmail.com", userPassword)

	t.Run("should get a user_usecase by id", func(t *testing.T) {
		repo := &UserRepository{}
		user, _ := repo.Insert(*userFix)

		got, err := repo.GetById(user.ID)

		assert.Nil(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, user.ID, got.ID)
		assert.Equal(t, user.Email.String(), got.Email.String())
		assert.True(t, got.Password.IsCorrectPassword(userPassword))
	})

	t.Run("should return a error when pass the name incorrect", func(t *testing.T) {
		repo := &UserRepository{}
		got, err := repo.GetById("")

		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.Equal(t, err, repository.ErrUserNotFound)
	})
}

func TestUserRepository_Insert(t *testing.T) {
	userPassword := "123456"
	userFix, _ := user.NewUser("Tiago", "tiagotigore@hotmail.com", userPassword)

	t.Run("should return a valid object when call the function insert", func(t *testing.T) {
		repo := &UserRepository{}

		got, err := repo.Insert(*userFix)

		assert.Nil(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, userFix.Name, got.Name)
		assert.Equal(t, userFix.Email, got.Email)
		assert.True(t, got.Password.IsCorrectPassword(userPassword))
		assert.NotNil(t, got.ID)
	})
}

func TestUserRepository_GetByEmail(t *testing.T) {
	userPassword := "123456"
	userFix, _ := user.NewUser("Tiago", "tiagotigore@hotmail.com", userPassword)

	t.Run("should return a valid object when call the function GetByEmail", func(t *testing.T) {
		repo := &UserRepository{}
		user, _ := repo.Insert(*userFix)
		got, err := repo.GetByEmail(user.Email.String())

		assert.Nil(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, user.Email, got.Email)
	})

	t.Run("should return a error when call the GetByEmail with object without email", func(t *testing.T) {
		repo := &UserRepository{}
		got, err := repo.GetByEmail("Email invalid")

		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.Equal(t, err, repository.ErrUserNotFound)
	})
}
