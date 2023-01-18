package task_usecase

import "todo_project.com/internal/app/repository"

type GetAllByClient interface {
	Handle(input GetAllByClientInput) (*[]Output, error)
}

type GetAllByClientInput struct {
	IDUser string
}

type getAllByClient struct {
	taskRepository repository.ITaskRepository
}

func NewGetAllByClient(taskRepository repository.ITaskRepository) GetAllByClient {
	return &getAllByClient{taskRepository: taskRepository}
}

func (g *getAllByClient) Handle(input GetAllByClientInput) (*[]Output, error) {
	tasks, err := g.taskRepository.GetAllByClientId(input.IDUser)
	if err != nil {
		return nil, err
	}

	return tasksOutputFromTasks(tasks), nil
}
