package task_usecase

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

type delete struct {
	taskRepository repository.ITaskRepository
}

func NewDelete(taskRepository repository.ITaskRepository) Delete {
	return &delete{taskRepository: taskRepository}
}

func (d *delete) Handle(input DeleteInput) (*DeleteOutput, error) {
	_, err := d.taskRepository.GetById(input.ID)
	if err != nil {
		return nil, err
	}

	err = d.taskRepository.DeleteById(input.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteOutput{}, nil
}
