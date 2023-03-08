package notes_usecase

import (
	"time"
	"todo_project.com/internal/domain/notes"
)

type Output struct {
	ID          string
	DateCreated *time.Time
	DateUpdate  *time.Time
	Title       string
	Description string
}

func notesOutputFromTask(notes *notes.Notes) *Output {
	return &Output{
		ID:          notes.ID,
		DateCreated: notes.DateCreated,
		DateUpdate:  notes.DateUpdate,
		Title:       notes.Title,
		Description: notes.Description,
	}
}

func notesOutputFromTasks(notes *[]notes.Notes) *[]Output {
	outputs := make([]Output, 0, len(*notes))
	for _, noteGot := range *notes {
		outputs = append(outputs, *notesOutputFromTask(&noteGot))
	}
	return &outputs
}
