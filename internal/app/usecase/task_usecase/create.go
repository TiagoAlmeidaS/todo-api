package task_usecase

import (
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/app/security"
	"todo_project.com/internal/domain/task"
)

type Create interface {
	Handle(input CreateInput) (*Output, error)
}

type create struct {
	taskRepository repository.ITaskRepository
}

type CreateInput struct {
	LoggedUser  security.User
	DateInit    string
	DateEnd     string
	Title       string
	Description string
	IDProject   string
}

func (c *create) Handle(input CreateInput) (*Output, error) {

	var dateInit, dateEnd time.Time
	if input.DateInit != "" {
		dateInit, _ = time.Parse(LayoutFromParseTime, input.DateInit)
	}

	if input.DateEnd != "" {
		dateEnd, _ = time.Parse(LayoutFromParseTime, input.DateEnd)
	}

	task, err := task.NewTask(input.LoggedUser.ID, dateInit, dateEnd, input.Title, input.Description, input.IDProject)
	if err != nil {
		return nil, err
	}

	task, err = c.taskRepository.Insert(*task)
	if err != nil {
		return nil, err
	}

	return taskOutputFromTask(task), nil
}

func NewCreate(taskRepository repository.ITaskRepository) Create {
	return &create{taskRepository: taskRepository}
}
