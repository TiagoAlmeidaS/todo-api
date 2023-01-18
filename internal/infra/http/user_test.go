package http

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	httpGO "net/http"
	"testing"
	"todo_project.com/internal/app/usecase/user_usecase"
	user2 "todo_project.com/internal/domain/user"
)

type mockUCRegister struct {
	mock.Mock
}

func (m *mockUCRegister) Handle(input user_usecase.RegisterInput) (*user_usecase.Output, error) {
	args := m.Called(input)
	var result *user_usecase.Output
	if args.Get(0) != nil {
		result = args.Get(0).(*user_usecase.Output)
	}

	return result, args.Error(1)
}

type mockUCLogin struct {
	mock.Mock
}

func (m *mockUCLogin) Handle(input user_usecase.LoginInput) (*user_usecase.Output, error) {
	args := m.Called(input)
	var result *user_usecase.Output
	if args.Get(0) != nil {
		result = args.Get(0).(*user_usecase.Output)
	}
	return result, args.Error(1)
}

func TestUserResponseFromUserOutput(t *testing.T) {
	t.Run("should return a user_usecase response from user_usecase output", func(t *testing.T) {
		output := user_usecase.Output{
			ID:    "id",
			Name:  "name",
			Email: "email@mail.com",
		}
		response := userResponseFromUserOutput(output)
		assert.Equal(t, output.Email, response.Email)
		assert.Equal(t, output.Name, response.Name)
		assert.Equal(t, output.ID, response.ID)
	})

	t.Run("should return an empty user_usecase response", func(t *testing.T) {
		response := userResponseFromUserOutput(user_usecase.Output{})
		assert.Equal(t, "", response.ID)
		assert.Equal(t, "", response.Name)
		assert.Equal(t, "", response.Email)
	})
}

func TestUserController_Register(t *testing.T) {
	name, email, password, token := "Tiago", "tiagotigore@hotmail.com", "123456", "TOKEN"
	output := &user_usecase.Output{
		ID:    "id",
		Name:  name,
		Email: email,
	}

	t.Run("should register a user_usecase", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"name\":\"%s\",\"email\":\"%s\",\"password\":\"%s\"}", name, email, password),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockUCRegister := new(mockUCRegister)
		mockUCRegister.On("Handle", mock.Anything).Return(output, nil)

		uc := UserController{Authenticator: mockAuth, UCRegister: mockUCRegister}

		response := uc.Register(request)

		assert.Equal(t, httpGO.StatusCreated, response.HttpCode)
		assert.Equal(t, "{\"name\":\"Tiago\",\"email\":\"tiagotigore@hotmail.com\",\"token\":\"TOKEN\"}", response.Body)
		mockUCRegister.AssertCalled(t, "Handle", mock.MatchedBy(func(i interface{}) bool {
			input := i.(user_usecase.RegisterInput)
			return name == input.Name && email == input.Email && password == input.Password
		}))
	})

	t.Run("shouldn't register a user_usecase with invalid JSON", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{name\":\"%s\",\"email\":\"%s\",\"password\":\"%s\"}", name, email, password),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockUCRegister := new(mockUCRegister)
		uc := UserController{Authenticator: mockAuth, UCRegister: mockUCRegister}

		response := uc.Register(request)
		assert.Equal(t, httpGO.StatusBadRequest, response.HttpCode)
		assert.Equal(t, "{\"message\":\"invalid json body\"}", response.Body)
	})

	t.Run("shouldn't register a user_usecase when use case return an error", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"name\":\"%s\",\"email\":\"invalid_mail\",\"password\":\"%s\"}", name, password),
		}
		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockUCRegister := new(mockUCRegister)
		mockUCRegister.On("Handle", mock.Anything).Return(output, user2.ErrEmailIsInvalid)

		uc := UserController{
			Authenticator: mockAuth, UCRegister: mockUCRegister,
		}

		response := uc.Register(request)
		assert.Equal(t, httpGO.StatusBadRequest, response.HttpCode)
		assert.Equal(t, "{\"message\":\"email is invalid\"}", response.Body)
	})

	t.Run("should't register a user_usecase when authenticator return an error", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"name\":\"%s\",\"email\":\"invalid_mail\",\"password\":\"%s\"}", name, password),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return("", errors.New("an error"))
		mockUCRegister := new(mockUCRegister)
		mockUCRegister.On("Handle", mock.Anything).Return(output, nil)

		uc := UserController{Authenticator: mockAuth, UCRegister: mockUCRegister}

		response := uc.Register(request)
		assert.Equal(t, httpGO.StatusInternalServerError, response.HttpCode)
		assert.Equal(t, "{\"message\":\"an error\"}", response.Body)
	})
}

func TestUserController_Login(t *testing.T) {
	name, email, password, token := "test", "test@mail.com", "password", "the_token"
	output := &user_usecase.Output{
		ID:    "id",
		Name:  name,
		Email: email,
	}

	t.Run("should login a user_usecase", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"email\":\"%s\",\"password\":\"%s\"}", email, password),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockUCLogin := new(mockUCLogin)
		mockUCLogin.On("Handle", mock.Anything).Return(output, nil)

		uc := UserController{
			Authenticator: mockAuth, UCLogin: mockUCLogin,
		}

		response := uc.Login(request)

		assert.Equal(t, httpGO.StatusOK, response.HttpCode)
		assert.Equal(t, fmt.Sprintf("{\"name\":\"%s\",\"email\":\"%s\",\"token\":\"the_token\"}", name, email), response.Body)
		mockUCLogin.AssertCalled(t, "Handle", mock.MatchedBy(func(i interface{}) bool {
			input := i.(user_usecase.LoginInput)
			return email == input.Email && password == input.Password
		}))
	})

	t.Run("shouldn't login a user_usecase with invalid JSON", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{email\":\"%s\",\"password\":\"%s\"}", email, password),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockUCLogin := new(mockUCLogin)
		mockUCLogin.On("Handle", mock.Anything).Return(output, nil)

		uc := UserController{Authenticator: mockAuth, UCLogin: mockUCLogin}

		response := uc.Login(request)

		assert.Equal(t, httpGO.StatusBadRequest, response.HttpCode)
	})

	t.Run("shouldn't login a user_usecase when use case return an error", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{email\":\"%s\",\"password\":\"%s\"}", email, password),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockUCLogin := new(mockUCLogin)
		mockUCLogin.On("Handle", mock.Anything).Return(output, user_usecase.ErrUserEmailPasswordWrong)

		uc := UserController{
			Authenticator: mockAuth, UCLogin: mockUCLogin,
		}

		response := uc.Login(request)

		assert.Equal(t, httpGO.StatusBadRequest, response.HttpCode)
	})

	t.Run("shoudn't login a user_usecase when authenticator return an error", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"name\":\"%s\",\"email\":\"invalid_mail\",\"password\":\"%s\"}", name, password),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return("", errors.New("an error"))
		mockUCLogin := new(mockUCLogin)
		mockUCLogin.On("Handle", mock.Anything).Return(output, nil)

		uc := UserController{
			Authenticator: mockAuth,
			UCLogin:       mockUCLogin,
		}

		response := uc.Login(request)

		assert.Equal(t, httpGO.StatusInternalServerError, response.HttpCode)
	})
}
