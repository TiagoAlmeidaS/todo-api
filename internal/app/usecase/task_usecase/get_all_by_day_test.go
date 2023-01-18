package task_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/domain/task"
	"todo_project.com/internal/infra/memory"
)

func TestGetAllByDay(t *testing.T) {
	dateTime := time.Now()
	taskFix, _ := task.NewTask("123", dateTime.Add(time.Hour*1), dateTime.Add(time.Hour*2), "task 1", "Estou escrevendo aqui a primera task_usecase", "")
	taskFix2, _ := task.NewTask("123", dateTime.Add(time.Hour*3), dateTime.Add(time.Hour*4), "task 2", "Estou escrevendo aqui a primera task_usecase", "")
	taskFix3, _ := task.NewTask("123", dateTime.Add(time.Hour*4), dateTime.Add(time.Hour*5), "task 3", "Estou escrevendo aqui a primera task_usecase", "")

	t.Run("should get the all tasks of client", func(t *testing.T) {
		repo := &memory.TaskRepository{}

		_, err := repo.Insert(*taskFix)
		_, err = repo.Insert(*taskFix2)
		_, err = repo.Insert(*taskFix2)
		_, err = repo.Insert(*taskFix3)

		input := GetAllByDayInput{
			Day:    dateTime.Add(time.Hour * 1).Add(time.Millisecond * 200),
			IDUser: taskFix.IDUser,
		}

		uc := NewGetAllByDay(repo)

		output, err := uc.Handle(input)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Len(t, *output, 1)
	})
}
