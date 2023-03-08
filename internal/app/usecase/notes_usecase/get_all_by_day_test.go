package notes_usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/domain/notes"
	"todo_project.com/internal/infra/memory"
)

func TestGetAllByDay(t *testing.T) {
	dateTime := time.Now()
	noteFix, _ := notes.NewNotes("123", dateTime.Add(time.Hour*1), dateTime.Add(time.Hour*2), "task 1", "Estou escrevendo aqui a primera task_usecase")
	noteFix2, _ := notes.NewNotes("123", dateTime.Add(time.Hour*3), dateTime.Add(time.Hour*4), "task 2", "Estou escrevendo aqui a primera task_usecase")
	noteFix3, _ := notes.NewNotes("123", dateTime.Add(time.Hour*25), dateTime.Add(time.Hour*5), "task 3", "Estou escrevendo aqui a primera task_usecase")

	t.Run("should get the all tasks of client", func(t *testing.T) {
		repo := &memory.NotesRepository{}

		_, err := repo.Insert(*noteFix)
		_, err = repo.Insert(*noteFix2)
		_, err = repo.Insert(*noteFix2)
		_, err = repo.Insert(*noteFix3)

		input := GetAllByDayInput{
			Day:    dateTime.Add(time.Hour * 1).Add(time.Millisecond * 200),
			IDUser: noteFix.IDUser,
		}

		uc := NewGetAllByDay(repo)

		output, err := uc.Handle(input)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Len(t, *output, 3)
	})
}
