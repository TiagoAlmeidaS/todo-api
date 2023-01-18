package task_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/domain/task"
	"todo_project.com/internal/infra/memory"
)

func TestGetResumeStatus(t *testing.T) {

	taskFix, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")
	taskFix2, _ := task.NewTask("12354", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")

	t.Run("should get resume status of client", func(t *testing.T) {

		repo := &memory.TaskRepository{}

		taskGot, _ := repo.Insert(*taskFix)
		taskGot, _ = repo.Insert(*taskFix)
		taskGot.Status = task.Process
		repo.Update(*taskGot)
		taskGot, _ = repo.Insert(*taskFix)
		taskGot, _ = repo.Insert(*taskFix)
		taskGot.Status = task.Completed
		repo.Update(*taskGot)
		repo.Insert(*taskFix2)

		output, err := repo.GetResumeStatus(taskGot.IDUser)
		assert.Nil(t, err)
		assert.Equal(t, output.Open, 2)
		assert.Equal(t, output.Completed, 1)
		assert.Equal(t, output.Process, 1)
	})
}
