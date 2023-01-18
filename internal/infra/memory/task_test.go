package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/domain/task"
)

func TestTaskRepository_GetById(t *testing.T) {
	taskFix, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")
	repo := &TaskRepository{}

	t.Run("should return a task_usecase with there id", func(t *testing.T) {

		taskGot, err := repo.Insert(*taskFix)
		assert.Nil(t, err)

		got, err := repo.GetById(taskGot.ID)
		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Equal(t, taskGot.ID, got.ID)
		assert.Equal(t, taskGot.IDUser, got.IDUser)
		assert.Equal(t, taskGot.Title, got.Title)
	})

	t.Run("should return a error when not found a task_usecase", func(t *testing.T) {
		got, err := repo.GetById("ID_INVALID")
		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.Equal(t, err, repository.ErrTaskNotFound)
	})
}

func TestTaskRepository_Insert(t *testing.T) {

	taskFix, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")

	t.Run("should insert a valid task_usecase", func(t *testing.T) {
		repo := &TaskRepository{}

		got, err := repo.Insert(*taskFix)

		assert.Nil(t, err)
		assert.NotNil(t, got)
		assert.NotNil(t, got.ID)
		assert.Equal(t, taskFix.Title, got.Title)
		assert.Equal(t, taskFix.IDUser, got.IDUser)
		assert.Equal(t, taskFix.Status, got.Status)
		assert.Equal(t, got.Status, task.Open)
		assert.Equal(t, taskFix.Description, got.Description)
		assert.Equal(t, taskFix.DateInit, got.DateInit)
		assert.Equal(t, taskFix.DateEnd, got.DateEnd)
	})
}

func TestTaskRepository_GetAllByClientId(t *testing.T) {
	taskFix, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")
	taskFix2, _ := task.NewTask("111", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")

	t.Run("should return a list of all tasks by client", func(t *testing.T) {
		repo := &TaskRepository{}

		_, err := repo.Insert(*taskFix)
		_, err = repo.Insert(*taskFix)
		_, err = repo.Insert(*taskFix2)
		_, err = repo.Insert(*taskFix2)

		assert.Nil(t, err)

		got, err := repo.GetAllByClientId("123")

		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Len(t, *got, 2)
	})

	t.Run("should return a list empty of tasjs when client don't have a insert", func(t *testing.T) {
		repo := &TaskRepository{}

		got, err := repo.GetAllByClientId("123")

		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Len(t, *got, 0)
	})
}

func TestTaskRepository_Update(t *testing.T) {
	taskFix, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")
	repo := &TaskRepository{}

	t.Run("should return a udpate task_usecase", func(t *testing.T) {
		taskGot, err := repo.Insert(*taskFix)
		assert.Nil(t, err)

		updateTitle := "Segundo Title"
		taskGot.Title = updateTitle
		taskUpdate, err := repo.Update(*taskGot)

		assert.Nil(t, err)
		assert.NotNil(t, taskUpdate)
		assert.Equal(t, taskUpdate.Title, updateTitle)
	})
}

func TestTaskRepository_DeleteById(t *testing.T) {
	taskFix, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")
	repo := &TaskRepository{}

	t.Run("should delete a task_usecase valid", func(t *testing.T) {

		taskGot, err := repo.Insert(*taskFix)
		assert.Nil(t, err)

		err = repo.DeleteById(taskGot.ID)
		assert.Nil(t, err)
	})
}

func TestTaskRepository_GetResumeStatus(t *testing.T) {
	taskFix, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")
	repo := &TaskRepository{}

	t.Run("should delete a task_usecase valid", func(t *testing.T) {

		taskGot, err := repo.Insert(*taskFix)
		taskGot, err = repo.Insert(*taskFix)
		taskGot.Status = task.Process
		_, err = repo.Update(*taskGot)
		taskGot, err = repo.Insert(*taskFix)
		taskGot, err = repo.Insert(*taskFix)
		assert.Nil(t, err)

		output, err := repo.GetResumeStatus(taskFix.IDUser)
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, output.Process, 1)
		assert.Equal(t, output.Completed, 0)
		assert.Equal(t, output.Open, 3)
	})
}

func TestTaskRepository_GetAllByDay(t *testing.T) {
	dateTime := time.Now()
	taskFix, _ := task.NewTask("123", dateTime.Add(time.Hour*1), dateTime.Add(time.Hour*2), "task 1", "Estou escrevendo aqui a primera task_usecase", "")
	taskFix2, _ := task.NewTask("123", dateTime.Add(time.Hour*3), dateTime.Add(time.Hour*4), "task 2", "Estou escrevendo aqui a primera task_usecase", "")
	taskFix3, _ := task.NewTask("123", dateTime.Add(time.Hour*4), dateTime.Add(time.Hour*5), "task 3", "Estou escrevendo aqui a primera task_usecase", "")
	repo := &TaskRepository{}

	t.Run("should get all task by day", func(t *testing.T) {
		_, err := repo.Insert(*taskFix)
		_, err = repo.Insert(*taskFix)
		_, err = repo.Insert(*taskFix2)
		_, err = repo.Insert(*taskFix2)
		_, err = repo.Insert(*taskFix2)
		_, err = repo.Insert(*taskFix3)

		var dateTimeAdd = dateTime.Add(time.Hour * 1).Add(time.Millisecond * 200)

		output, err := repo.GetAllByDay(dateTimeAdd, "123")
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Len(t, *output, 2)

		dateTimeAdd = dateTime.Add(time.Hour * 3).Add(time.Millisecond * 200)
		output, err = repo.GetAllByDay(dateTimeAdd, "123")
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Len(t, *output, 3)

		dateTimeAdd = dateTime.Add(time.Hour * 4).Add(time.Millisecond * 200)
		output, err = repo.GetAllByDay(dateTimeAdd, "123")
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Len(t, *output, 1)
	})
}
