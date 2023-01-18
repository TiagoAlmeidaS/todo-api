package task_usecase

import (
	"time"
	"todo_project.com/internal/app/repository"
)

type GetAllByDayInput struct {
	Day    time.Time
	IDUser string
}

type GetAllByDay interface {
	Handle(input GetAllByDayInput) (*[]Output, error)
}

type getAllByDay struct {
	taskRepository repository.ITaskRepository
}

func NewGetAllByDay(taskRepository repository.ITaskRepository) GetAllByDay {
	return &getAllByDay{taskRepository: taskRepository}
}

func (g *getAllByDay) Handle(input GetAllByDayInput) (*[]Output, error) {
	tasks, err := g.taskRepository.GetAllByDay(input.Day, input.IDUser)
	if err != nil {
		return nil, err
	}

	return tasksOutputFromTasks(tasks), nil
}
