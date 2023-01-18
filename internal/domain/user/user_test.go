package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {

	name, email, pass := "Tiago Almeida dos Santos", "teste123@gmail.com", "123456"

	t.Run("should return a new user_usecase correctly", func(t *testing.T) {
		newUser, err := NewUser(name, email, pass)

		assert.Nil(t, err)
		assert.Equal(t, name, newUser.Name)
		assert.Equal(t, email, newUser.Email.String())
		assert.True(t, newUser.Password.IsCorrectPassword(pass))
	})

	t.Run("should return a error when user_usecase name is empty", func(t *testing.T) {
		newUser, err := NewUser("", email, pass)

		assert.Nil(t, newUser)
		assert.Equal(t, err, ErrNameIsInvalid)
	})

	t.Run("should return a error when put a invalid email", func(t *testing.T) {
		newUser, err := NewUser(name, "tiago", pass)

		assert.Nil(t, newUser)
		assert.Equal(t, err, ErrEmailIsInvalid)

		newUser, err = NewUser(name, "", pass)

		assert.Nil(t, newUser)
		assert.Equal(t, err, ErrEmailIsInvalid)

		newUser, err = NewUser(name, "tiago@", pass)

		assert.Nil(t, newUser)
		assert.Equal(t, err, ErrEmailIsInvalid)
	})

	t.Run("should return error when put a invalid password", func(t *testing.T) {
		newUser, err := NewUser(name, email, "")

		assert.Nil(t, newUser)
		assert.Equal(t, err, ErrPasswordIsInvalid)

		newUser, err = NewUser(name, email, "123")

		assert.Nil(t, newUser)
		assert.Equal(t, err, ErrPasswordIsInvalid)
	})
}
