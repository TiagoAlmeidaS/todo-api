package task

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {

	iDUser, dateInit, dateEnd, title, description := "123", time.Now(), time.Now(), "Toma", "Tomaaa"

	t.Run("Should return a New Task", func(t *testing.T) {
		response, err := NewTask(iDUser, dateInit, dateEnd, title, description, "123")
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, response.IDUser, iDUser)
		assert.Equal(t, *response.DateInit, dateInit)
		assert.Equal(t, *response.DateEnd, dateEnd)
		assert.Equal(t, response.Title, title)
		assert.Equal(t, response.Description, description)
	})

	t.Run("should return a error when call NewTask invalid title", func(t *testing.T) {
		response, err := NewTask(iDUser, dateInit, dateEnd, "", description, "123")
		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, err, ErrTitleIsInvalid)
	})
}

func TestSetProcessing(t *testing.T) {
	iDUser, dateInit, dateEnd, title, description := "123", time.Now(), time.Now(), "Toma", "Tomaaa"
	task, _ := NewTask(iDUser, dateInit, dateEnd, title, description, "123")

	t.Run("Should return a valid processing", func(t *testing.T) {
		err := task.SetProcessing()
		assert.Nil(t, err)
		assert.Equal(t, task.Status, Process)
	})
}

func TestClose(t *testing.T) {
	iDUser, dateInit, dateEnd, title, description := "123", time.Now(), time.Now(), "Toma", "Tomaaa"
	task, _ := NewTask(iDUser, dateInit, dateEnd, title, description, "")
	t.Run("should return a valid close task_usecase", func(t *testing.T) {
		err := task.Close()
		assert.Nil(t, err)
		assert.Equal(t, task.Status, Completed)
	})
}
