package task_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/domain/task"
	"todo_project.com/internal/infra/memory"
)

func TestNewGetAllByClient(t *testing.T) {
	taskFix, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")
	taskFix2, _ := task.NewTask("1423", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")
	t.Run("should get the all tasks of client", func(t *testing.T) {
		repo := &memory.TaskRepository{}

		taskGot, _ := repo.Insert(*taskFix)
		taskGot, _ = repo.Insert(*taskFix)
		taskGot, _ = repo.Insert(*taskFix)
		taskGot, _ = repo.Insert(*taskFix)
		_, _ = repo.Insert(*taskFix2)

		input := GetAllByClientInput{
			IDUser: taskGot.IDUser,
		}

		uc := NewGetAllByClient(repo)

		output, err := uc.Handle(input)
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Len(t, *output, 4)
	})
}
