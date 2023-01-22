package notes_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/domain/notes"
	"todo_project.com/internal/infra/memory"
)

func TestDelete(t *testing.T) {
	noteFix, _ := notes.NewNotes("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase")

	t.Run("should delete the task_usecase", func(t *testing.T) {
		repo := &memory.NotesRepository{}

		task, _ := repo.Insert(*noteFix)

		uc := NewDelete(repo)

		input := DeleteInput{
			ID: task.ID,
		}

		output, err := uc.Handle(input)
		assert.Nil(t, err)
		assert.NotNil(t, output)
	})

	t.Run("should return a error when don't locale the task_usecase", func(t *testing.T) {
		repo := &memory.NotesRepository{}

		uc := NewDelete(repo)

		input := DeleteInput{
			ID: "ID_INVALID",
		}

		output, err := uc.Handle(input)
		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.NotNil(t, repository.ErrTaskNotFound)

	})
}
