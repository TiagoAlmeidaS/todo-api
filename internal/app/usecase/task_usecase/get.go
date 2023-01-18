package task_usecase

import (
	"todo_project.com/internal/app/repository"
)

type Get interface {
	Handle(input GetInput) (*Output, error)
}

type get struct {
	taskRepository repository.ITaskRepository
}

type GetInput struct {
	IDTask string
}

func (g get) Handle(input GetInput) (*Output, error) {
	task, err := g.taskRepository.GetById(input.IDTask)
	if err != nil {
		return nil, err
	}

	return &Output{
		ID:          task.ID,
		DateInit:    task.DateInit,
		DateEnd:     task.DateEnd,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		IDProject:   task.IDProject,
	}, nil
}

func NewGet(taskRepository repository.ITaskRepository) Get {
	return &get{taskRepository: taskRepository}
}
