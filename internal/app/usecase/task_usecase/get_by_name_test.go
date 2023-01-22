package task_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/domain/task"
	"todo_project.com/internal/infra/memory"
)

func TestNewGetByName(t *testing.T) {
	taskFix, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")
	taskFix2, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "Segunda task_usecase tiago", "Estou escrevendo aqui a primera task_usecase", "")
	taskFix4, _ := task.NewTask("12345", time.Now(), time.Now().Add(time.Hour*5), "Segunda task_usecase tiago", "Estou escrevendo aqui a primera task_usecase", "")
	taskFix3, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "toma", "Estou escrevendo aqui a primera task_usecase", "")
	t.Run("should get the filter tasks of client", func(t *testing.T) {
		repo := &memory.TaskRepository{}

		_, _ = repo.Insert(*taskFix)
		_, _ = repo.Insert(*taskFix)
		_, _ = repo.Insert(*taskFix)
		_, _ = repo.Insert(*taskFix3)
		_, _ = repo.Insert(*taskFix2)
		_, _ = repo.Insert(*taskFix2)
		_, _ = repo.Insert(*taskFix4)

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
