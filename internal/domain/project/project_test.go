package project

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/domain/task"
)

func TestProject(t *testing.T) {
	var tasks []task.Task
	idUser, title, description, dateEnd, dateUpdate := "123", "Novo project", "nova descricao", time.Now(), time.Now()
	t.Run("should create a new project", func(t *testing.T) {
		output, err := NewProject(idUser, dateEnd, dateUpdate, title, description, tasks)
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, output.Title, title)
	})

	t.Run("should return a error when the title is empty", func(t *testing.T) {
		output, err := NewProject(idUser, dateEnd, dateUpdate, "", description, tasks)
		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, err, ErrTitleIsInvalid)
	})

	t.Run("should return a error when the title is empty", func(t *testing.T) {
		output, err := NewProject("", dateEnd, dateUpdate, title, description, tasks)
		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, err, ErrIdUserIsInvalid)
	})
}
