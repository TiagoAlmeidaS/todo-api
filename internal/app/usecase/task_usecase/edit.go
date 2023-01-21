package task_usecase

import (
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/domain/task"
)

const (
	LayoutFromParseTime = "2006-01-02T15:04:05"
)

type Edit interface {
	Handle(input EditInput) (*EditOutput, error)
}

type edit struct {
	taskRepository repository.ITaskRepository
}

type EditInput struct {
	ID          string
	IDUser      string
	DateInit    string
	DateEnd     string
	Title       string
	Description string
	Status      task.Status
	IDProject   string
}

type EditOutput struct {
	DateInit    *time.Time
	DateEnd     *time.Time
	Title       string
	Description string
	Status      task.Status
	IDProject   string
}

func NewEdit(taskRepository repository.ITaskRepository) Edit {
	return &edit{taskRepository: taskRepository}
}

func (e *edit) Handle(input EditInput) (*EditOutput, error) {
	taskGot, err := e.taskRepository.GetById(input.ID)
	if err != nil {
		return nil, err
	}

	if input.IDUser != taskGot.IDUser {
		return nil, err
	}

	var dateInit, dateEnd time.Time
	if input.DateInit != "" {
		dateInit, _ = time.Parse(LayoutFromParseTime, input.DateInit)
	}

	if input.DateEnd != "" {
		dateEnd, _ = time.Parse(LayoutFromParseTime, input.DateEnd)
	}

	taskGot.DateInit = &dateInit
	taskGot.DateEnd = &dateEnd
	taskGot.Title = input.Title
	taskGot.Description = input.Description
	taskGot.Status = input.Status
	taskGot.IDProject = input.IDProject

	taskGot, err = e.taskRepository.Update(*taskGot)
	if err != nil {
		return nil, err
	}

	return &EditOutput{
		DateInit:    taskGot.DateInit,
		DateEnd:     taskGot.DateEnd,
		Title:       taskGot.Title,
		Description: taskGot.Description,
		Status:      taskGot.Status,
	}, nil
}
