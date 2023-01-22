package notes_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/app/security"
	"todo_project.com/internal/domain/notes"
	"todo_project.com/internal/domain/task"
	"todo_project.com/internal/domain/user"
	"todo_project.com/internal/infra/memory"
)

func TestCreate(t *testing.T) {
	noteFix, _ := notes.NewNotes("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase")

	t.Run("should return a new Task with status Open", func(t *testing.T) {
		repo := &memory.NotesRepository{}
		repoUser := &memory.UserRepository{}

		uc := NewCreate(repo)

		userRegister, _ := user.NewUser("tiago", "email@email.com", "123456")

		userGot, _ := repoUser.Insert(*userRegister)

		intput := CreateInput{
			Title:       noteFix.Title,
			Description: noteFix.Description,
			LoggedUser:  security.NewUser(*userGot),
		}

		output, err := uc.Handle(intput)
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.NotEmpty(t, output.ID)
		assert.Equal(t, output.Title, noteFix.Title)
	})

	t.Run("should return a error when object is incorrect", func(t *testing.T) {
		repo := &memory.NotesRepository{}

		uc := NewCreate(repo)

		input := CreateInput{
			Title:       noteFix.Title,
			Description: noteFix.Description,
			LoggedUser:  security.NewUser(user.User{}),
		}

		output, err := uc.Handle(input)
		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, err, task.ErrIdUserIsInvalid)
	})
}
