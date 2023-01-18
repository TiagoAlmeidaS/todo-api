package repository

type AllRepositories struct {
	IUserRepository IUserRepository
	ITaskRepository ITaskRepository
}

type Repositories interface {
	GetRepositories() (*AllRepositories, error)
}
