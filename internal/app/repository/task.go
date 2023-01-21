package repository

import (
	"errors"
	"time"
	"todo_project.com/internal/domain/task"
)

var (
	ErrTaskNotFound       = errors.New("task not found")
	ErrTaskUpdateNotFound = errors.New("task update not found")
)

type ITaskRepository interface {
	GetById(id string) (*task.Task, error)
	Insert(task task.Task) (*task.Task, error)
	Update(task task.Task) (*task.Task, error)
	GetAllByClientId(clientId string) (*[]task.Task, error)
	DeleteById(id string) error
	GetResumeStatus(clientId string) (*task.Resume, error)
	GetAllByDay(day time.Time, clientId string) (*[]task.Task, error)
}
