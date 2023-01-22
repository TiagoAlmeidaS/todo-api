package notes_usecase

import (
	"todo_project.com/internal/app/repository"
)

type Get interface {
	Handle(input GetInput) (*Output, error)
}

type get struct {
	notesRepository repository.INotesRepository
}

type GetInput struct {
	IDTask string
	IDUser string
}

func (g get) Handle(input GetInput) (*Output, error) {
	note, err := g.notesRepository.GetById(input.IDTask)
	if err != nil {
		return nil, err
	}

	return &Output{
		ID:          note.ID,
		DateCreated: note.DateCreated,
		DateUpdate:  note.DateUpdate,
		Title:       note.Title,
		Description: note.Description,
	}, nil
}

func NewGet(notesRepository repository.INotesRepository) Get {
	return &get{notesRepository: notesRepository}
}
