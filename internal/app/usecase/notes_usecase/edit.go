package notes_usecase

import (
	"time"
	"todo_project.com/internal/app/repository"
)

type Edit interface {
	Handle(input EditInput) (*Output, error)
}

type edit struct {
	notesRepository repository.INotesRepository
}

type EditInput struct {
	ID          string
	IDUser      string
	Title       string
	Description string
	DateUpdate  time.Time
}

func NewEdit(notesRepository repository.INotesRepository) Edit {
	return &edit{notesRepository: notesRepository}
}

func (e *edit) Handle(input EditInput) (*Output, error) {
	notesGot, err := e.notesRepository.GetById(input.ID)
	if err != nil {
		return nil, err
	}

	if input.IDUser != notesGot.IDUser {
		return nil, err
	}

	notesGot.DateUpdate = &input.DateUpdate
	notesGot.Title = input.Title
	notesGot.Description = input.Description

	notesGot, err = e.notesRepository.Update(*notesGot)
	if err != nil {
		return nil, err
	}

	return &Output{
		DateUpdate:  notesGot.DateUpdate,
		DateCreated: notesGot.DateCreated,
		ID:          notesGot.ID,
		Title:       notesGot.Title,
		Description: notesGot.Description,
	}, nil
}
