package task_usecase

import "todo_project.com/internal/app/repository"

type GetResumeStatus interface {
	Handle(input GetResumeStatusInput) (*OutputResume, error)
}

type GetResumeStatusInput struct {
	IDUser string
}

type getResumeStatus struct {
	taskRepository repository.ITaskRepository
}

func NewGetResumeStatus(taskRepository repository.ITaskRepository) GetResumeStatus {
	return &getResumeStatus{taskRepository: taskRepository}
}

func (g *getResumeStatus) Handle(input GetResumeStatusInput) (*OutputResume, error) {
	resumeGot, err := g.taskRepository.GetResumeStatus(input.IDUser)
	if err != nil {
		return nil, err
	}

	return resumeOutputFromTaskResume(resumeGot), nil
}