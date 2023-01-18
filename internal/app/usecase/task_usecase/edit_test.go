package task_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/domain/task"
	"todo_project.com/internal/infra/memory"
)

func TestEdit(t *testing.T) {

	taskFix, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")

	t.Run("should edit task_usecase", func(t *testing.T) {
		repo := &memory.TaskRepository{}

		_, err := repo.Insert(*taskFix)
		taskGot, err := repo.Insert(*taskFix)
		_, err = repo.Insert(*taskFix)

		newTitle := "Novo title"
		newDescription := "Nova descricao"

		uc := NewEdit(repo)
		input := EditInput{
			ID:          taskGot.ID,
			Status:      task.Process,
			Description: newDescription,
			Title:       newTitle,
			DateInit:    taskGot.DateInit.String(),
		}

		output, err := uc.Handle(input)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, output.Title, input.Title)
		assert.Equal(t, task.Process, output.Status)
	})

}
