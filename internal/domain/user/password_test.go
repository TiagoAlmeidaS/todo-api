package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPassword(t *testing.T) {
	t.Run("should return a new password valid", func(t *testing.T) {
		passExpect := "123456"
		validPass := true
		receive, err := NewPassword(passExpect)

		assert.Nil(t, err)
		assert.NotNil(t, receive)
		assert.Equal(t, validPass, receive.IsCorrectPassword(passExpect))
	})

	t.Run("should return a error when call the function with a password less than 6", func(t *testing.T) {
		pass := "1234"
		receive, err := NewPassword(pass)

		assert.Nil(t, receive)
		assert.NotNil(t, err)
		assert.Equal(t, ErrPasswordIsInvalid, err)
	})

}

func TestIsCorrectPassword(t *testing.T) {
	t.Run("should return a valid password", func(t *testing.T) {
		expectPassword := "123456"
		receive, err := NewPassword(expectPassword)
		assert.Nil(t, err)
		assert.NotNil(t, receive)
		assert.True(t, receive.IsCorrectPassword(expectPassword))
	})

	t.Run("should return a invalid password when password is diferent", func(t *testing.T) {
		expectPass := "123456"
		receive, err := NewPassword(expectPass)
		assert.Nil(t, err)
		assert.False(t, receive.IsCorrectPassword("654321"))
	})
}

func TestNewPasswordHashed(t *testing.T) {
	t.Run("should return a new hash when call the method NewPasswordHashed", func(t *testing.T) {
		expectPass := "123456"
		receive, err := NewPasswordHashed(expectPass)
		assert.Nil(t, err)
		assert.NotNil(t, receive)
		assert.Equal(t, expectPass, receive.String())
	})
}
