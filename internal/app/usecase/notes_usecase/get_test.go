package notes_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/domain/notes"
	"todo_project.com/internal/domain/user"
	"todo_project.com/internal/infra/memory"
)

func TestGet(t *testing.T) {

	noteFix, _ := notes.NewNotes("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase")

	t.Run("should get the note", func(t *testing.T) {
		repo := &memory.NotesRepository{}
		repoUser := &memory.UserRepository{}

		userRegister, _ := user.NewUser("tiago", "email@email.com", "123456")

		userGot, _ := repoUser.Insert(*userRegister)

		noteGot, _ := repo.Insert(*noteFix)

		uc := NewGet(repo)

		input := GetInput{
			IDTask: noteGot.ID,
			IDUser: userGot.ID,
		}

		output, err := uc.Handle(input)
		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, output.ID, noteGot.ID)
		assert.Equal(t, output.Title, noteGot.Title)
		assert.Equal(t, output.Description, noteGot.Description)
	})

	t.Run("should return a error when don't exists a notes valid", func(t *testing.T) {
		repo := &memory.NotesRepository{}
		repoUser := &memory.UserRepository{}
		userRegister, _ := user.NewUser("tiago", "email@email.com", "123456")
		userGot, _ := repoUser.Insert(*userRegister)

		uc := NewGet(repo)
		input := GetInput{
			IDTask: "ID_TASK_INVALID",
			IDUser: userGot.ID,
		}

		output, err := uc.Handle(input)
		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, err, repository.ErrTaskNotFound)
	})
}
