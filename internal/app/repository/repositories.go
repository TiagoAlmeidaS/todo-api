package repository

type AllRepositories struct {
	IUserRepository  IUserRepository
	ITaskRepository  ITaskRepository
	INotesRepository INotesRepository
}

type Repositories interface {
	GetRepositories() (*AllRepositories, error)
}
