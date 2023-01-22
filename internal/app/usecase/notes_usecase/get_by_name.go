package notes_usecase

import (
	"todo_project.com/internal/app/repository"
)

type GetByNameInput struct {
	NameTask string
	IDUser   string
}

type GetByName interface {
	Handle(input GetByNameInput) (*[]Output, error)
}

type getByName struct {
	notesRepository repository.INotesRepository
}

func NewGetByName(notesRepository repository.INotesRepository) GetByName {
	return &getByName{notesRepository: notesRepository}
}

func (g *getByName) Handle(input GetByNameInput) (*[]Output, error) {
	tasks, err := g.notesRepository.GetByName(input.NameTask, input.IDUser)
	if err != nil {
		return nil, err
	}

	return notesOutputFromTasks(tasks), nil
}
