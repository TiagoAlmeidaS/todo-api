package task_usecase

import (
	"time"
	"todo_project.com/internal/domain/task"
)

type Status string

type Output struct {
	ID          string
	DateInit    *time.Time
	DateEnd     *time.Time
	Title       string
	Description string
	Status      task.Status
	IDProject   string
}

func taskOutputFromTask(task *task.Task) *Output {
	return &Output{
		ID:          task.ID,
		DateInit:    task.DateInit,
		DateEnd:     task.DateEnd,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		IDProject:   task.IDProject,
	}
}

func tasksOutputFromTasks(tasks *[]task.Task) *[]Output {
	outputs := make([]Output, 0, len(*tasks))
	for _, taskGot := range *tasks {
		outputs = append(outputs, *taskOutputFromTask(&taskGot))
	}
	return &outputs
}

type OutputResume struct {
	Open      int
	Process   int
	Completed int
}

func resumeOutputFromTaskResume(resume *task.Resume) *OutputResume {
	return &OutputResume{
		Open:      resume.Open,
		Process:   resume.Process,
		Completed: resume.Completed,
	}
}
