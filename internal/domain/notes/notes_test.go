package notes

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewNotes(t *testing.T) {
	title, description, dateCreated, dateUpdate, idUser := "Primeiro teste", "descricao do primeiro teste", time.Now(), time.Now(), "123"
	t.Run("should new notes", func(t *testing.T) {

		newNote, err := NewNotes(idUser, dateCreated, dateUpdate, title, description)

		assert.Nil(t, err)
		assert.NotNil(t, newNote)
		assert.Equal(t, newNote.Title, title)
	})

	t.Run("should return a erro when notes is incorrectly", func(t *testing.T) {
		newNote, err := NewNotes(idUser, dateCreated, dateUpdate, "", description)
		assert.Nil(t, newNote)
		assert.NotNil(t, err)
		assert.Equal(t, err, ErrTitleIsInvalid)
	})

}
