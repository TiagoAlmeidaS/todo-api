package repository

import (
	"errors"
	"time"
	"todo_project.com/internal/domain/notes"
)

var (
	ErrNotesNotFound       = errors.New("note not found")
	ErrNotesUpdateNotFound = errors.New("note update not found")
)

type INotesRepository interface {
	GetById(id string) (*notes.Notes, error)
	Insert(notes notes.Notes) (*notes.Notes, error)
	Update(notes notes.Notes) (*notes.Notes, error)
	GetAllByClientId(clientId string) (*[]notes.Notes, error)
	DeleteById(id string) error
	GetAllByDay(day time.Time, clientId string) (*[]notes.Notes, error)
	GetByName(nameNotes string, clientId string) (*[]notes.Notes, error)
}
