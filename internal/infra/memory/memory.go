package memory

import "todo_project.com/internal/app/repository"

type Repositories struct{}

func (r *Repositories) GetRepositories() (*repository.AllRepositories, error) {
	return &repository.AllRepositories{
		IUserRepository: &UserRepository{},
		ITaskRepository: &TaskRepository{},
	}, nil
}
