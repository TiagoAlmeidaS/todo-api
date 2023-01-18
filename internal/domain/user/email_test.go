package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {

	t.Run("should return a new and valid email", func(t *testing.T) {
		expectEmail := "teste@gmail.com"
		receive, err := NewEmail(expectEmail)
		assert.Nil(t, err)
		assert.NotNil(t, receive)
		assert.Equal(t, expectEmail, receive.String())
	})

	t.Run("should return a error when email were invalid", func(t *testing.T) {
		emailError := "testegmail.com"
		receive, err := NewEmail(emailError)

		assert.Nil(t, receive)
		assert.NotNil(t, err)
		assert.Equal(t, err, ErrEmailIsInvalid)
	})

	t.Run("should return a error when email is empty", func(t *testing.T) {
		emailEmpty := ""
		receive, err := NewEmail(emailEmpty)

		assert.Nil(t, receive)
		assert.NotNil(t, err)
		assert.Equal(t, err, ErrEmailIsInvalid)
	})

}
