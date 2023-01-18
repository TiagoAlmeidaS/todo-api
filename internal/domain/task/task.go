package task

import (
	"errors"
	"time"
)

type Status string

var (
	ErrTitleIsInvalid  = errors.New("title is invalid")
	ErrIdUserIsInvalid = errors.New("user_usecase is not valid")
	ErrDateIsInvalid   = errors.New("data is not valid")
)

const (
	LayoutFromParseTime        = "2006-01-02T15:04:05"
	Open                Status = "OPEN"
	Process             Status = "PROCESS"
	Completed           Status = "COMPLETED"
)

type Resume struct {
	Open      int
	Process   int
	Completed int
}

type Task struct {
	ID          string
	IDUser      string
	DateInit    *time.Time
	DateEnd     *time.Time
	Title       string
	Description string
	Status      Status
	IDProject   string
}

func NewTask(IDUser string, dataInit time.Time, dataEnd time.Time, title string, description string, idProject string) (*Task, error) {
	if title == "" {
		return nil, ErrTitleIsInvalid
	}
	if IDUser == "" {
		return nil, ErrIdUserIsInvalid
	}
	if &dataInit == nil {
		return nil, ErrDateIsInvalid
	}

	return &Task{
		IDUser:      IDUser,
		DateInit:    &dataInit,
		DateEnd:     &dataEnd,
		Status:      Open,
		Description: description,
		Title:       title,
		IDProject:   idProject,
	}, nil
}

func (e *Task) SetProcessing() error {
	e.Status = Process
	return nil
}

func (e *Task) Close() error {
	e.Status = Completed
	return nil
}
