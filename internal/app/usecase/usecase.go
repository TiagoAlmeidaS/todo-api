package usecase

import (
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/app/usecase/task_usecase"
	"todo_project.com/internal/app/usecase/user_usecase"
)

type AllUserCases struct {
	UserLogin           user_usecase.Login
	UserRegister        user_usecase.Register
	TaskCreate          task_usecase.Create
	TaskDelete          task_usecase.Delete
	TaskEdit            task_usecase.Edit
	TaskGet             task_usecase.Get
	TaskGetAllByClient  task_usecase.GetAllByClient
	TaskGetAllByDay     task_usecase.GetAllByDay
	TaskGetResumeStatus task_usecase.GetResumeStatus
}

func GetUseCases(repositories repository.Repositories) (*AllUserCases, error) {
	repos, err := repositories.GetRepositories()
	if err != nil {
		return nil, err
	}

	return &AllUserCases{
		UserLogin:           user_usecase.NewLogin(repos.IUserRepository),
		UserRegister:        user_usecase.NewRegister(repos.IUserRepository),
		TaskCreate:          task_usecase.NewCreate(repos.ITaskRepository),
		TaskEdit:            task_usecase.NewEdit(repos.ITaskRepository),
		TaskDelete:          task_usecase.NewDelete(repos.ITaskRepository),
		TaskGet:             task_usecase.NewGet(repos.ITaskRepository),
		TaskGetAllByClient:  task_usecase.NewGetAllByClient(repos.ITaskRepository),
		TaskGetAllByDay:     task_usecase.NewGetAllByDay(repos.ITaskRepository),
		TaskGetResumeStatus: task_usecase.NewGetResumeStatus(repos.ITaskRepository),
	}, nil
}
