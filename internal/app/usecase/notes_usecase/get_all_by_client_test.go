package notes_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/domain/notes"
	"todo_project.com/internal/infra/memory"
)

func TestNewGetAllByClient(t *testing.T) {
	notesFix, _ := notes.NewNotes("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase")
	taskFix2, _ := notes.NewNotes("1423", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase")
	t.Run("should get the all tasks of client", func(t *testing.T) {
		repo := &memory.NotesRepository{}

		noteGot, _ := repo.Insert(*notesFix)
		noteGot, _ = repo.Insert(*notesFix)
		noteGot, _ = repo.Insert(*notesFix)
		noteGot, _ = repo.Insert(*notesFix)
		_, _ = repo.Insert(*taskFix2)

		input := GetAllByClientInput{
			IDUser: noteGot.IDUser,
		}

		uc := NewGetAllByClient(repo)

		output, err := uc.Handle(input)
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Len(t, *output, 4)
	})
}
