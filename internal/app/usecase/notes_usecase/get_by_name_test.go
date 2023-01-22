package notes_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/domain/notes"
	"todo_project.com/internal/infra/memory"
)

func TestNewGetByName(t *testing.T) {
	noteFix, _ := notes.NewNotes("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase")
	noteFix2, _ := notes.NewNotes("123", time.Now(), time.Now().Add(time.Hour*5), "Segunda task_usecase tiago", "Estou escrevendo aqui a primera task_usecase")
	noteFix4, _ := notes.NewNotes("12345", time.Now(), time.Now().Add(time.Hour*5), "Segunda task_usecase tiago", "Estou escrevendo aqui a primera task_usecase")
	noteFix3, _ := notes.NewNotes("123", time.Now(), time.Now().Add(time.Hour*5), "toma", "Estou escrevendo aqui a primera task_usecase")
	t.Run("should get the filter tasks of client", func(t *testing.T) {
		repo := &memory.NotesRepository{}

		_, _ = repo.Insert(*noteFix)
		_, _ = repo.Insert(*noteFix)
		_, _ = repo.Insert(*noteFix)
		_, _ = repo.Insert(*noteFix3)
		_, _ = repo.Insert(*noteFix2)
		_, _ = repo.Insert(*noteFix2)
		_, _ = repo.Insert(*noteFix4)

		input := GetByNameInput{
			NameTask: "ti",
			IDUser:   "123",
		}

		uc := NewGetByName(repo)

		output, err := uc.Handle(input)
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Len(t, *output, 2)
	})
}
