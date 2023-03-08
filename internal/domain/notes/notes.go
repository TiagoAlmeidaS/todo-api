package notes

import (
	"errors"
	"time"
)

var (
	ErrTitleIsInvalid  = errors.New("title is invalid")
	ErrIdUserIsInvalid = errors.New("user_usecase is not valid")
)

type Notes struct {
	ID          string
	IDUser      string
	DateCreated *time.Time
	DateUpdate  *time.Time
	Title       string
	Description string
}

func NewNotes(IDUser string, dateCreated time.Time, dateUpdate time.Time, title string, description string) (*Notes, error) {
	if title == "" {
		return nil, ErrTitleIsInvalid
	}
	if IDUser == "" {
		return nil, ErrIdUserIsInvalid
	}

	return &Notes{
		IDUser:      IDUser,
		DateCreated: &dateCreated,
		DateUpdate:  &dateUpdate,
		Description: description,
		Title:       title,
	}, nil
}
