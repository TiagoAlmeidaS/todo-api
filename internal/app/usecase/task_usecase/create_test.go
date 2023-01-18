package task_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/app/security"
	"todo_project.com/internal/domain/task"
	"todo_project.com/internal/domain/user"
	"todo_project.com/internal/infra/memory"
)

func TestCreate(t *testing.T) {
	taskFix, _ := task.NewTask("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase", "")

	t.Run("should return a new Task with status Open", func(t *testing.T) {
		repo := &memory.TaskRepository{}
		repoUser := &memory.UserRepository{}

		uc := NewCreate(repo)

		userRegister, _ := user.NewUser("tiago", "email@email.com", "123456")

		user, _ := repoUser.Insert(*userRegister)

		intput := CreateInput{
			Title:       taskFix.Title,
			Description: taskFix.Description,
			LoggedUser:  security.NewUser(*user),
			DateInit:    taskFix.DateInit.String(),
			DateEnd:     taskFix.DateEnd.String(),
		}

		output, err := uc.Handle(intput)
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.NotEmpty(t, output.ID)
		assert.Equal(t, output.Title, taskFix.Title)
	})

	t.Run("should return a error when object is incorrect", func(t *testing.T) {
		repo := &memory.TaskRepository{}

		uc := NewCreate(repo)

		input := CreateInput{
			Title:       taskFix.Title,
			Description: taskFix.Description,
			LoggedUser:  security.NewUser(user.User{}),
			DateInit:    taskFix.DateInit.String(),
			DateEnd:     taskFix.DateEnd.String(),
		}

		output, err := uc.Handle(input)
		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, err, task.ErrIdUserIsInvalid)
	})
}
