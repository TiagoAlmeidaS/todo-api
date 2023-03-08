package notes_usecase

import (
	"time"
	"todo_project.com/internal/app/repository"
)

type GetAllByDayInput struct {
	Day    time.Time
	IDUser string
}

type GetAllByDay interface {
	Handle(input GetAllByDayInput) (*[]Output, error)
}

type getAllByDay struct {
	notesRepository repository.INotesRepository
}

func NewGetAllByDay(notesRepository repository.INotesRepository) GetAllByDay {
	return &getAllByDay{notesRepository: notesRepository}
}

func (g *getAllByDay) Handle(input GetAllByDayInput) (*[]Output, error) {
	tasks, err := g.notesRepository.GetAllByDay(input.Day, input.IDUser)
	if err != nil {
		return nil, err
	}

	return notesOutputFromTasks(tasks), nil
}
