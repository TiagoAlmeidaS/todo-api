package notes_usecase

import (
	"todo_project.com/internal/app/repository"
)

type Delete interface {
	Handle(input DeleteInput) (*DeleteOutput, error)
}

type DeleteInput struct {
	ID string
}

type DeleteOutput struct{}

type deleteNotes struct {
	notesRepository repository.INotesRepository
}

func NewDelete(notesRepository repository.INotesRepository) Delete {
	return &deleteNotes{notesRepository: notesRepository}
}

func (d *deleteNotes) Handle(input DeleteInput) (*DeleteOutput, error) {
	_, err := d.notesRepository.GetById(input.ID)
	if err != nil {
		return nil, err
	}

	err = d.notesRepository.DeleteById(input.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{}, nil
}
