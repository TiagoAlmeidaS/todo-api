package task_usecase

import (
	"todo_project.com/internal/app/repository"
)

type GetByNameInput struct {
	NameTask string
	IDUser   string
}

type GetByName interface {
	Handle(input GetByNameInput) (*[]Output, error)
}

type getByName struct {
	taskRepository repository.ITaskRepository
}

func NewGetByName(taskRepository repository.ITaskRepository) GetByName {
	return &getByName{taskRepository: taskRepository}
}

func (g *getByName) Handle(input GetByNameInput) (*[]Output, error) {
	tasks, err := g.taskRepository.GetByName(input.NameTask, input.IDUser)
	if err != nil {
		return nil, err
	}

	return tasksOutputFromTasks(tasks), nil
}
