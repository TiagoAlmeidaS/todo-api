package project

import (
	"errors"
	"time"
	"todo_project.com/internal/domain/task"
)

var (
	ErrTitleIsInvalid  = errors.New("title is invalid")
	ErrIdUserIsInvalid = errors.New("user_usecase is not valid")
)

type Project struct {
	ID          string
	IDUser      string
	DateEnd     *time.Time
	DateUpdate  *time.Time
	Title       string
	Description string
	tasks       []task.Task
}

func NewProject(idUser string, dateEnd time.Time, dateUpdate time.Time, title string, description string, tasks []task.Task) (*Project, error) {
	if idUser == "" {
		return nil, ErrIdUserIsInvalid
	}

	if title == "" {
		return nil, ErrTitleIsInvalid
	}

	return &Project{
		IDUser:      idUser,
		DateEnd:     &dateEnd,
		DateUpdate:  &dateUpdate,
		Title:       title,
		Description: description,
		tasks:       tasks,
	}, nil
}
