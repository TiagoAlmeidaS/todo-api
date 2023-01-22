package http

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	httpGO "net/http"
	"testing"
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/app/security"
	"todo_project.com/internal/app/usecase/notes_usecase"
	"todo_project.com/internal/domain/task"
)

type mockNCCreate struct {
	mock.Mock
}

func (m *mockNCCreate) Handle(input notes_usecase.CreateInput) (*notes_usecase.Output, error) {
	args := m.Called(input)
	var result *notes_usecase.Output
	if args.Get(0) != nil {
		result = args.Get(0).(*notes_usecase.Output)
	}
	return result, args.Error(1)
}

type mockNCGetByName struct {
	mock.Mock
}

func (m *mockNCGetByName) Handle(input notes_usecase.GetByNameInput) (*[]notes_usecase.Output, error) {
	args := m.Called(input)
	var result *[]notes_usecase.Output
	if args.Get(0) != nil {
		result = args.Get(0).(*[]notes_usecase.Output)
	}
	return result, args.Error(1)
}

type mockNCGet struct {
	mock.Mock
}

func (m *mockNCGet) Handle(input notes_usecase.GetInput) (*notes_usecase.Output, error) {
	args := m.Called(input)
	var result *notes_usecase.Output
	if args.Get(0) != nil {
		result = args.Get(0).(*notes_usecase.Output)
	}
	return result, args.Error(1)
}

type mockNCGetAllByClientId struct {
	mock.Mock
}

func (m *mockNCGetAllByClientId) Handle(input notes_usecase.GetAllByClientInput) (*[]notes_usecase.Output, error) {
	args := m.Called(input)
	var result *[]notes_usecase.Output
	if args.Get(0) != nil {
		result = args.Get(0).(*[]notes_usecase.Output)
	}
	return result, args.Error(1)
}

type mockNCDelete struct {
	mock.Mock
}

func (m *mockNCDelete) Handle(input notes_usecase.DeleteInput) (*notes_usecase.DeleteOutput, error) {
	args := m.Called(input)
	var result *notes_usecase.DeleteOutput
	if args.Get(0) != nil {
		result = args.Get(0).(*notes_usecase.DeleteOutput)
	}
	return result, args.Error(1)
}

type mockNCEdit struct {
	mock.Mock
}

func (m *mockNCEdit) Handle(input notes_usecase.EditInput) (*notes_usecase.Output, error) {
	args := m.Called(input)
	var result *notes_usecase.Output
	if args.Get(0) != nil {
		result = args.Get(0).(*notes_usecase.Output)
	}
	return result, args.Error(1)
}

type mockNCGetAllByDay struct {
	mock.Mock
}

func (m *mockNCGetAllByDay) Handle(input notes_usecase.GetAllByDayInput) (*[]notes_usecase.Output, error) {
	args := m.Called(input)
	var result *[]notes_usecase.Output
	if args.Get(0) != nil {
		result = args.Get(0).(*[]notes_usecase.Output)
	}
	return result, args.Error(1)
}

func TestNotesResponseFromOutput(t *testing.T) {
	t.Run("should return a notes_usecase response from output", func(t *testing.T) {
		dateString := "2019-01-01T00:00:00"
		date, _ := time.Parse("2006-01-02T15:04:05", dateString)
		output := notes_usecase.Output{
			ID:          "id",
			Title:       "title",
			Description: "description",
			DateCreated: &date,
			DateUpdate:  &date,
		}

		got := notesResponseFromOutput(output)

		assert.NotNil(t, got)
		assert.Equal(t, output.Title, got.Title)
	})

	t.Run("should return a empty notes_usecase", func(t *testing.T) {
		got := notesResponseFromOutput(notes_usecase.Output{})

		assert.Equal(t, got.Title, "")
	})
}

func TestNotesResponseFromOutputs(t *testing.T) {
	t.Run("should return a list of tasks", func(t *testing.T) {
		got := notesResponseFromOutputs([]notes_usecase.Output{{}, {}})

		assert.Len(t, *got, 2)
	})
}

func TestNotesController_Create(t *testing.T) {
	date := time.Now()

	title, description, token := "title", "description", "the_token"

	output := &notes_usecase.Output{
		ID:          "id",
		Title:       "title",
		Description: "description",
		DateCreated: &date,
		DateUpdate:  &date,
	}

	t.Run("should register a notes_usecase", func(t *testing.T) {

		request := Request{
			LoggedUser: &security.User{},
			Body:       fmt.Sprintf("{\"date_init\":\"%s\",\"date_end\":\"%s\",\"title\":\"%s\",\"description\":\"%s\"}", date, date, title, description),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockNCCreate := new(mockNCCreate)
		mockNCCreate.On("Handle", mock.Anything).Return(output, nil)

		uc := NotesController{Authenticator: mockAuth, NCCreate: mockNCCreate}

		response := uc.Create(request)

		assert.NotNil(t, response)
		assert.Equal(t, response.HttpCode, httpGO.StatusCreated)
	})

	t.Run("should return a error a notes_usecase when object is invalid", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\":\"%s\",\"date_init\":\"%s\",\"date_end\":\"%s\",\"title\":\"%s\",\"description\":\"%s\"}", "", date, date, title, description),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockNCCreate := new(mockNCCreate)
		mockNCCreate.On("Handle", mock.Anything).Return(nil, task.ErrIdUserIsInvalid)

		uc := NotesController{Authenticator: mockAuth, NCCreate: mockNCCreate}
		response := uc.Create(request)

		assert.NotNil(t, response)
		assert.Equal(t, response.HttpCode, httpGO.StatusInternalServerError)
	})
}

func TestNotesController_Get(t *testing.T) {
	date := time.Now()

	title, description, token := "title", "description", "the_token"

	output := &notes_usecase.Output{
		ID:          "id",
		Title:       title,
		Description: description,
		DateCreated: &date,
		DateUpdate:  &date,
	}

	t.Run("should get a notes_usecase valid", func(t *testing.T) {

		request := Request{
			Body: fmt.Sprintf("{\"id\":\"%s\"}", "123"),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockNCGet := new(mockNCGet)
		mockNCGet.On("Handle", mock.Anything).Return(output, nil)

		uc := NotesController{Authenticator: mockAuth, NCGet: mockNCGet}

		response := uc.Get(request)

		assert.NotNil(t, response)
		assert.Equal(t, response.HttpCode, httpGO.StatusOK)
	})

	t.Run("should return error when don't exists notes_usecase", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"id\":\"%s\"}", "123"),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockNCGet := new(mockNCGet)
		mockNCGet.On("Handle", mock.Anything).Return(nil, repository.ErrTaskNotFound)

		uc := NotesController{Authenticator: mockAuth, NCGet: mockNCGet}

		response := uc.Get(request)

		assert.NotNil(t, response)
		assert.Equal(t, response.HttpCode, httpGO.StatusNotFound)
	})
}

func TestNotesController_GetAllByClientId(t *testing.T) {

	t.Run("should get all note by client", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"id_user\":\"%s\"}", "123"),
		}
		date := time.Now()

		token := "the_token"

		output := &[]notes_usecase.Output{{
			ID:          "id",
			Title:       "title",
			Description: "description",
			DateCreated: &date,
			DateUpdate:  &date,
		},
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockNCGetAllByClientId := new(mockNCGetAllByClientId)
		mockNCGetAllByClientId.On("Handle", mock.Anything).Return(output, nil)

		uc := NotesController{Authenticator: mockAuth, NCGetAllByClient: mockNCGetAllByClientId}
		response := uc.GetAllByClientId(request)

		assert.Equal(t, response.HttpCode, httpGO.StatusOK)
	})

	t.Run("should return empty note", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"id_user\":\"%s\"}", "123"),
		}

		token := "the_token"

		output := &[]notes_usecase.Output{{}}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockNCGetAllByClientId := new(mockNCGetAllByClientId)
		mockNCGetAllByClientId.On("Handle", mock.Anything).Return(output, nil)

		uc := NotesController{Authenticator: mockAuth, NCGetAllByClient: mockNCGetAllByClientId}
		response := uc.GetAllByClientId(request)

		assert.Equal(t, response.HttpCode, httpGO.StatusOK)
	})

}

func TestNotesController_Delete(t *testing.T) {

	token := "the_token"

	t.Run("should delete the note", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"id\":\"%s\"}", "123"),
		}

		output := &notes_usecase.DeleteOutput{}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockNCDelete := new(mockNCDelete)
		mockNCDelete.On("Handle", mock.Anything).Return(output, nil)

		uc := NotesController{Authenticator: mockAuth, NCDelete: mockNCDelete}

		response := uc.Delete(request)

		assert.NotNil(t, response)
		assert.Equal(t, response.HttpCode, httpGO.StatusOK)
	})

	t.Run("should return a error when note not found", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"id\":\"%s\"}", "123"),
		}

		//output := &notes_usecase.DeleteOutput{}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockNCDelete := new(mockNCDelete)
		mockNCDelete.On("Handle", mock.Anything).Return(nil, repository.ErrTaskNotFound)

		uc := NotesController{Authenticator: mockAuth, NCDelete: mockNCDelete}

		response := uc.Delete(request)

		assert.NotNil(t, response)
		assert.Equal(t, response.HttpCode, httpGO.StatusInternalServerError)
	})
}

func TestNotesController_Edit(t *testing.T) {
	date := time.Now()

	title, token, description2 := "title", "the_token", "description2"

	output := &notes_usecase.Output{
		Title:       title,
		Description: description2,
		DateCreated: &date,
		DateUpdate:  &date,
		ID:          "123",
	}

	t.Run("should edit a note", func(t *testing.T) {

		params := map[string]string{"id": "123"}

		request := Request{
			Params:     params,
			Body:       fmt.Sprintf("{\"date_init\":\"%s\",\"date_end\":\"%s\",\"title\":\"%s\",\"description\":\"%s\", \"IDProject\":\"\"}", date, date, title, description2),
			LoggedUser: &security.User{},
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockNCEdit := new(mockNCEdit)
		mockNCEdit.On("Handle", mock.Anything).Return(output, nil)

		uc := NotesController{Authenticator: mockAuth, NCEdit: mockNCEdit}

		response := uc.Edit(request)

		assert.Equal(t, response.HttpCode, httpGO.StatusOK)
	})
}

func TestNotesController_GetAllByDay(t *testing.T) {
	params := map[string]string{"day": "2006-01-02T15:04:05"}

	t.Run("should get all task by day", func(t *testing.T) {
		request := Request{
			Params:     params,
			LoggedUser: &security.User{ID: "123"},
		}

		date := time.Now()

		token := "the_token"

		output := &[]notes_usecase.Output{{
			ID:          "id",
			Title:       "title",
			Description: "description",
			DateCreated: &date,
			DateUpdate:  &date,
		},
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockNCGetAllByDay := new(mockNCGetAllByDay)
		mockNCGetAllByDay.On("Handle", mock.Anything).Return(output, nil)

		uc := NotesController{Authenticator: mockAuth, NCGetAllByDay: mockNCGetAllByDay}
		response := uc.GetAllByDay(request)

		assert.Equal(t, httpGO.StatusOK, response.HttpCode)
	})

}

func TestNotesController_GetByName(t *testing.T) {
	params := map[string]string{"name": "ti"}

	t.Run("should get all notes by day", func(t *testing.T) {
		request := Request{
			Params:     params,
			LoggedUser: &security.User{ID: "123"},
		}

		date := time.Now()

		token := "the_token"

		output := &[]notes_usecase.Output{{
			ID:          "id",
			Title:       "title",
			Description: "description",
			DateCreated: &date,
			DateUpdate:  &date,
		},
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockNCGetByName := new(mockNCGetByName)
		mockNCGetByName.On("Handle", mock.Anything).Return(output, nil)

		uc := NotesController{Authenticator: mockAuth, NCGetByName: mockNCGetByName}
		response := uc.GetByName(request)

		assert.Equal(t, httpGO.StatusOK, response.HttpCode)
	})

}
