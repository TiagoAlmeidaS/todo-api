package http

import (
	"encoding/json"
	"net/http"
	"todo_project.com/internal/app/security"
	"todo_project.com/internal/app/usecase/user_usecase"
	user2 "todo_project.com/internal/domain/user"
)

type (
	UserRegisterRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	UserLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	UserResponse struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Token string `json:"token"`
	}
)

func userResponseFromUserOutput(user user_usecase.Output) *UserResponse {
	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

type UserController struct {
	UCRegister    user_usecase.Register
	UCLogin       user_usecase.Login
	Authenticator security.Authenticator
}

func (c *UserController) Register(request Request) Response {
	var registerRequest UserRegisterRequest
	err := json.Unmarshal([]byte(request.Body), &registerRequest)
	if err != nil {
		return Response{
			HttpCode: http.StatusBadRequest,
			Body:     wrapError(ErrInvalidJsonBody),
		}
	}

	input := user_usecase.RegisterInput{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}

	output, err := c.UCRegister.Handle(input)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		switch err {
		case user_usecase.ErrUserAlreadyExists,
			user2.ErrNameIsInvalid,
			user2.ErrEmailIsInvalid,
			user2.ErrPasswordIsInvalid:
			httpStatus = http.StatusBadRequest
		}
		return Response{
			HttpCode: httpStatus,
			Body:     wrapError(err),
		}
	}
	token, err := c.Authenticator.Generate(security.User{
		ID:   output.ID,
		Name: output.Name,
	})
	if err != nil {
		return Response{
			HttpCode: http.StatusInternalServerError,
			Body:     wrapError(err),
		}
	}
	return Response{
		HttpCode: http.StatusCreated,
		Body: wrapBody(UserResponse{
			Token: token,
			Email: output.Email,
			Name:  output.Name,
			ID:    output.ID,
		}),
	}
}

func (c *UserController) Login(request Request) Response {
	var loginRequest UserLoginRequest
	err := json.Unmarshal([]byte(request.Body), &loginRequest)
	if err != nil {
		return Response{
			HttpCode: http.StatusBadRequest,
			Body:     wrapError(err),
		}
	}

	input := user_usecase.LoginInput{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}

	output, err := c.UCLogin.Handle(input)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		switch err {
		case user_usecase.ErrUserEmailPasswordWrong:
			httpStatus = http.StatusBadRequest
		}
		return Response{
			HttpCode: httpStatus,
			Body:     wrapError(err),
		}
	}

	token, err := c.Authenticator.Generate(security.User{
		ID:   output.ID,
		Name: output.Name,
	})
	if err != nil {
		return Response{
			HttpCode: http.StatusInternalServerError,
			Body:     wrapError(err),
		}
	}

	return Response{
		HttpCode: http.StatusOK,
		Body: wrapBody(UserResponse{
			Token: token,
			Email: output.Email,
			Name:  output.Name,
			ID:    output.ID,
		}),
	}
}
