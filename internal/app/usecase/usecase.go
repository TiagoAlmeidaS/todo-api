package usecase

import (
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/app/usecase/notes_usecase"
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
	TaskGetByName       task_usecase.GetByName
	NotesCreate         notes_usecase.Create
	NotesDelete         notes_usecase.Delete
	NotesEdit           notes_usecase.Edit
	NotesGet            notes_usecase.Get
	NotesGetAllByClient notes_usecase.GetAllByClient
	NotesGetAllByDay    notes_usecase.GetAllByDay
	NotesGetByName      notes_usecase.GetByName
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
		TaskGetByName:       task_usecase.NewGetByName(repos.ITaskRepository),
		NotesCreate:         notes_usecase.NewCreate(repos.INotesRepository),
		NotesEdit:           notes_usecase.NewEdit(repos.INotesRepository),
		NotesDelete:         notes_usecase.NewDelete(repos.INotesRepository),
		NotesGet:            notes_usecase.NewGet(repos.INotesRepository),
		NotesGetAllByClient: notes_usecase.NewGetAllByClient(repos.INotesRepository),
		NotesGetAllByDay:    notes_usecase.NewGetAllByDay(repos.INotesRepository),
		NotesGetByName:      notes_usecase.NewGetByName(repos.INotesRepository),
	}, nil
}
