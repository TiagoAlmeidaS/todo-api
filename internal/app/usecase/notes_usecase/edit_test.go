package notes_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/domain/notes"
	"todo_project.com/internal/infra/memory"
)

func TestEdit(t *testing.T) {

	notesFix, _ := notes.NewNotes("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase")

	t.Run("should edit task_usecase", func(t *testing.T) {
		repo := &memory.NotesRepository{}

		_, err := repo.Insert(*notesFix)
		taskGot, err := repo.Insert(*notesFix)
		_, err = repo.Insert(*notesFix)

		newTitle := "Novo title"
		newDescription := "Nova descricao"

		uc := NewEdit(repo)
		input := EditInput{
			ID:          taskGot.ID,
			IDUser:      "123",
			Description: newDescription,
			Title:       newTitle,
			DateUpdate:  time.Now().Add(time.Hour * 5),
		}

		output, err := uc.Handle(input)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, output.Title, input.Title)
	})

}
