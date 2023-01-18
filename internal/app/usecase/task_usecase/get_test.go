package task_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/domain/task"
	"todo_project.com/internal/infra/memory"
)

func TestGet(t *testing.T) {

	taskFix, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")

	t.Run("should get the task_usecase", func(t *testing.T) {
		repo := &memory.TaskRepository{}

		taskGot, _ := repo.Insert(*taskFix)

		uc := NewGet(repo)

		input := GetInput{
			IDTask: taskGot.ID,
		}

		output, err := uc.Handle(input)
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, output.ID, taskGot.ID)
		assert.Equal(t, output.Title, taskGot.Title)
		assert.Equal(t, output.Status, taskGot.Status)
		assert.Equal(t, output.Description, taskGot.Description)
	})

	t.Run("should return a error when don't exists a task_usecase valid", func(t *testing.T) {
		repo := &memory.TaskRepository{}

		uc := NewGet(repo)
		input := GetInput{
			IDTask: "ID_TASK_INVALID",
		}

		output, err := uc.Handle(input)
		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, err, repository.ErrTaskNotFound)
	})
}
