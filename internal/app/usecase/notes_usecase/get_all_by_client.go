package notes_usecase

import "todo_project.com/internal/app/repository"

type GetAllByClient interface {
	Handle(input GetAllByClientInput) (*[]Output, error)
}

type GetAllByClientInput struct {
	IDUser string
}

type getAllByClient struct {
	notesRepository repository.INotesRepository
}

func NewGetAllByClient(notesRepository repository.INotesRepository) GetAllByClient {
	return &getAllByClient{notesRepository: notesRepository}
}

func (g *getAllByClient) Handle(input GetAllByClientInput) (*[]Output, error) {
	tasks, err := g.notesRepository.GetAllByClientId(input.IDUser)
	if err != nil {
		return nil, err
	}

	return notesOutputFromTasks(tasks), nil
}
