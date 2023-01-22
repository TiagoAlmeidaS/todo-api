package notes_usecase

import (
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/app/security"
	"todo_project.com/internal/domain/notes"
)

type Create interface {
	Handle(input CreateInput) (*Output, error)
}

type create struct {
	notesRepository repository.INotesRepository
}

type CreateInput struct {
	LoggedUser  security.User
	Title       string
	Description string
}

func (c *create) Handle(input CreateInput) (*Output, error) {

	note, err := notes.NewNotes(input.LoggedUser.ID, time.Now(), time.Now(), input.Title, input.Description)
	if err != nil {
		return nil, err
	}

	note, err = c.notesRepository.Insert(*note)
	if err != nil {
		return nil, err
	}

	return notesOutputFromTask(note), nil
}

func NewCreate(notesRepository repository.INotesRepository) Create {
	return &create{notesRepository: notesRepository}
}
