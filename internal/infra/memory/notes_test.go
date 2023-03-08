package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/domain/notes"
)

func TestNotesRepository_GetById(t *testing.T) {
	notesFix, _ := notes.NewNotes("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase")
	repo := &NotesRepository{}

	t.Run("should return a task_usecase with there id", func(t *testing.T) {

		notesGot, err := repo.Insert(*notesFix)
		assert.Nil(t, err)

		got, err := repo.GetById(notesGot.ID)
		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Equal(t, notesGot.ID, got.ID)
		assert.Equal(t, notesGot.IDUser, got.IDUser)
		assert.Equal(t, notesGot.Title, got.Title)
	})

	t.Run("should return a error when not found a task_usecase", func(t *testing.T) {
		got, err := repo.GetById("ID_INVALID")
		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.Equal(t, err, repository.ErrTaskNotFound)
	})
}

func TestNotesRepository_Insert(t *testing.T) {

	notesFix, _ := notes.NewNotes("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase")

	t.Run("should insert a valid task_usecase", func(t *testing.T) {
		repo := &NotesRepository{}

		got, err := repo.Insert(*notesFix)

		assert.Nil(t, err)
		assert.NotNil(t, got)
		assert.NotNil(t, got.ID)
		assert.Equal(t, notesFix.Title, got.Title)
		assert.Equal(t, notesFix.IDUser, got.IDUser)
		assert.Equal(t, notesFix.Description, got.Description)
		assert.Equal(t, notesFix.DateUpdate, got.DateUpdate)
		assert.Equal(t, notesFix.DateCreated, got.DateCreated)
	})
}

func TestNotesRepository_GetAllByClientId(t *testing.T) {
	notesFix, _ := notes.NewNotes("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase")
	notesFix2, _ := notes.NewNotes("111", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase")

	t.Run("should return a list of all tasks by client", func(t *testing.T) {
		repo := &NotesRepository{}

		_, err := repo.Insert(*notesFix)
		_, err = repo.Insert(*notesFix)
		_, err = repo.Insert(*notesFix2)
		_, err = repo.Insert(*notesFix2)

		assert.Nil(t, err)

		got, err := repo.GetAllByClientId("123")

		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Len(t, *got, 2)
	})

	t.Run("should return a list empty of tasjs when client don't have a insert", func(t *testing.T) {
		repo := &NotesRepository{}

		got, err := repo.GetAllByClientId("123")

		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Len(t, *got, 0)
	})
}

func TestNotesRepository_Update(t *testing.T) {
	notesFix, _ := notes.NewNotes("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase")
	repo := &NotesRepository{}

	t.Run("should return a udpate task_usecase", func(t *testing.T) {
		notesGot, err := repo.Insert(*notesFix)
		assert.Nil(t, err)

		updateTitle := "Segundo Title"
		notesGot.Title = updateTitle
		taskUpdate, err := repo.Update(*notesGot)

		assert.Nil(t, err)
		assert.NotNil(t, taskUpdate)
		assert.Equal(t, taskUpdate.Title, updateTitle)
	})
}

func TestNotesRepository_DeleteById(t *testing.T) {
	notesFix, _ := notes.NewNotes("123", time.Now(), time.Now().Add(time.Hour*5), "Primeira task_usecase", "Estou escrevendo aqui a primera task_usecase")
	repo := &NotesRepository{}

	t.Run("should delete a task_usecase valid", func(t *testing.T) {

		notesGot, err := repo.Insert(*notesFix)
		assert.Nil(t, err)

		err = repo.DeleteById(notesGot.ID)
		assert.Nil(t, err)
	})
}

func TestNotesRepository_GetAllByDay(t *testing.T) {
	dateTime := time.Now()
	notesFix, _ := notes.NewNotes("123", dateTime.Add(time.Hour*1), dateTime.Add(time.Hour*2), "task 1", "Estou escrevendo aqui a primera task_usecase")
	notesFix2, _ := notes.NewNotes("123", dateTime.Add(time.Hour*3), dateTime.Add(time.Hour*4), "task 2", "Estou escrevendo aqui a primera task_usecase")
	notesFix3, _ := notes.NewNotes("123", dateTime.Add(time.Hour*4), dateTime.Add(time.Hour*5), "task 3", "Estou escrevendo aqui a primera task_usecase")
	repo := &NotesRepository{}

	t.Run("should get all task by day", func(t *testing.T) {
		_, err := repo.Insert(*notesFix)
		_, err = repo.Insert(*notesFix)
		_, err = repo.Insert(*notesFix2)
		_, err = repo.Insert(*notesFix2)
		_, err = repo.Insert(*notesFix2)
		_, err = repo.Insert(*notesFix3)

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
