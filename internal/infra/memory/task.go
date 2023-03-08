package memory

import (
	"github.com/google/uuid"
	"strings"
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/domain/task"
)

type TaskRepository struct {
	tasks []task.Task
}

func (r *TaskRepository) GetById(id string) (*task.Task, error) {
	for _, taskIn := range r.tasks {
		if taskIn.ID == id {
			return &taskIn, nil
		}
	}

	return nil, repository.ErrTaskNotFound
}

func (r *TaskRepository) Insert(task task.Task) (*task.Task, error) {
	task.ID = uuid.New().String()

	r.tasks = append(r.tasks, task)
	return &task, nil
}

func (r *TaskRepository) Update(task task.Task) (*task.Task, error) {
	for i, t := range r.tasks {
		if t.ID == task.ID {
			r.tasks[i] = task
			return &task, nil
		}
	}

	return nil, repository.ErrTaskNotFound
}

func (r *TaskRepository) GetAllByClientId(IDUser string) (*[]task.Task, error) {
	var tasksCreated []task.Task

	for _, taskIn := range r.tasks {
		if taskIn.IDUser == IDUser {
			tasksCreated = append(tasksCreated, taskIn)
		}
	}

	return &tasksCreated, nil
}

func (r *TaskRepository) GetResumeStatus(clientId string) (*task.Resume, error) {
	var open, process, completed int

	for _, taskFor := range r.tasks {
		if taskFor.IDUser == clientId {
			if taskFor.Status == task.Open {
				open = open + 1
			}

			if taskFor.Status == task.Process {
				process = process + 1
			}

			if taskFor.Status == task.Completed {
				completed = completed + 1
			}
		}
	}

	return &task.Resume{
		Open:      open,
		Process:   process,
		Completed: completed,
	}, nil
}

func (r *TaskRepository) DeleteById(id string) error {
	for i, taskIn := range r.tasks {
		if taskIn.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}

	return repository.ErrTaskNotFound
}

func (r *TaskRepository) GetAllByDay(day time.Time, clientId string) (*[]task.Task, error) {
	var taskFilted []task.Task

	for _, taskIn := range r.tasks {
		if taskIn.IDUser == clientId {
			if taskIn.DateEnd.After(day) && taskIn.DateInit.Before(day) {
				taskFilted = append(taskFilted, taskIn)
			}
		}
	}

	return &taskFilted, nil
}

func (r *TaskRepository) GetByName(nameTask string, idUser string) (*[]task.Task, error) {
	var taskFilted []task.Task

	for _, taskIn := range r.tasks {
		if strings.Contains(strings.ToLower(taskIn.Title), strings.ToLower(nameTask)) && taskIn.IDUser == idUser {
			taskFilted = append(taskFilted, taskIn)
		}
	}

	return &taskFilted, nil
}
