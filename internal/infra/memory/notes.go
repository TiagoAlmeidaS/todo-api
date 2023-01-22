package memory

import (
	"github.com/google/uuid"
	"strings"
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/domain/notes"
)

type NotesRepository struct {
	notes []notes.Notes
}

func (r *NotesRepository) GetById(id string) (*notes.Notes, error) {
	for _, noteIn := range r.notes {
		if noteIn.ID == id {
			return &noteIn, nil
		}
	}

	return nil, repository.ErrTaskNotFound
}

func (r *NotesRepository) Insert(task notes.Notes) (*notes.Notes, error) {
	task.ID = uuid.New().String()

	r.notes = append(r.notes, task)
	return &task, nil
}

func (r *NotesRepository) Update(task notes.Notes) (*notes.Notes, error) {
	for i, t := range r.notes {
		if t.ID == task.ID {
			r.notes[i] = task
			return &task, nil
		}
	}

	return nil, repository.ErrTaskNotFound
}

func (r *NotesRepository) GetAllByClientId(IDUser string) (*[]notes.Notes, error) {
	var tasksCreated []notes.Notes

	for _, noteIn := range r.notes {
		if noteIn.IDUser == IDUser {
			tasksCreated = append(tasksCreated, noteIn)
		}
	}

	return &tasksCreated, nil
}

func (r *NotesRepository) DeleteById(id string) error {
	for i, noteIn := range r.notes {
		if noteIn.ID == id {
			r.notes = append(r.notes[:i], r.notes[i+1:]...)
			return nil
		}
	}

	return repository.ErrTaskNotFound
}

func (r *NotesRepository) GetAllByDay(dateInput time.Time, clientId string) (*[]notes.Notes, error) {
	var notesFilted []notes.Notes

	for _, noteIn := range r.notes {
		if noteIn.IDUser == clientId {
			year, month, day := noteIn.DateCreated.Date()
			yearInput, monthInput, dayInput := dateInput.Date()
			if (year == yearInput) && (month == monthInput) && (day == dayInput) {
				notesFilted = append(notesFilted, noteIn)
			}
		}
	}

	return &notesFilted, nil
}

func (r *NotesRepository) GetByName(nameTask string, idUser string) (*[]notes.Notes, error) {
	var notesFilted []notes.Notes

	for _, noteIn := range r.notes {
		if strings.Contains(strings.ToLower(noteIn.Title), strings.ToLower(nameTask)) && noteIn.IDUser == idUser {
			notesFilted = append(notesFilted, noteIn)
		}
	}

	return &notesFilted, nil
}
