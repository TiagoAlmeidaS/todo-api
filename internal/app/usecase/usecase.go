package usecase

import (
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/app/usecase/user_usecase"
)

type AllUserCases struct {
	UserLogin    user_usecase.Login
	UserRegister user_usecase.Register
}

func GetUseCases(repositories repository.Repositories) (*AllUserCases, error) {
	repos, err := repositories.GetRepositories()
	if err != nil {
		return nil, err
	}

	return &AllUserCases{
		UserLogin:    user_usecase.NewLogin(repos.IUserRepository),
		UserRegister: user_usecase.NewRegister(repos.IUserRepository),
	}, nil
}
