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
	"todo_project.com/internal/app/usecase/task_usecase"
	"todo_project.com/internal/domain/task"
)

type mockTCCreate struct {
	mock.Mock
}

func (m *mockTCCreate) Handle(input task_usecase.CreateInput) (*task_usecase.Output, error) {
	args := m.Called(input)
	var result *task_usecase.Output
	if args.Get(0) != nil {
		result = args.Get(0).(*task_usecase.Output)
	}
	return result, args.Error(1)
}

type mockTCGet struct {
	mock.Mock
}

func (m *mockTCGet) Handle(input task_usecase.GetInput) (*task_usecase.Output, error) {
	args := m.Called(input)
	var result *task_usecase.Output
	if args.Get(0) != nil {
		result = args.Get(0).(*task_usecase.Output)
	}
	return result, args.Error(1)
}

type mockTCGetAllByClientId struct {
	mock.Mock
}

func (m *mockTCGetAllByClientId) Handle(input task_usecase.GetAllByClientInput) (*[]task_usecase.Output, error) {
	args := m.Called(input)
	var result *[]task_usecase.Output
	if args.Get(0) != nil {
		result = args.Get(0).(*[]task_usecase.Output)
	}
	return result, args.Error(1)
}

type mockTCDelete struct {
	mock.Mock
}

func (m *mockTCDelete) Handle(input task_usecase.DeleteInput) (*task_usecase.DeleteOutput, error) {
	args := m.Called(input)
	var result *task_usecase.DeleteOutput
	if args.Get(0) != nil {
		result = args.Get(0).(*task_usecase.DeleteOutput)
	}
	return result, args.Error(1)
}

type mockTCEdit struct {
	mock.Mock
}

func (m *mockTCEdit) Handle(input task_usecase.EditInput) (*task_usecase.EditOutput, error) {
	args := m.Called(input)
	var result *task_usecase.EditOutput
	if args.Get(0) != nil {
		result = args.Get(0).(*task_usecase.EditOutput)
	}
	return result, args.Error(1)
}

type mockTCGetAllByDay struct {
	mock.Mock
}

func (m *mockTCGetAllByDay) Handle(input task_usecase.GetAllByDayInput) (*[]task_usecase.Output, error) {
	args := m.Called(input)
	var result *[]task_usecase.Output
	if args.Get(0) != nil {
		result = args.Get(0).(*[]task_usecase.Output)
	}
	return result, args.Error(1)
}

type mockTCGetResumeStatus struct {
	mock.Mock
}

func (m *mockTCGetResumeStatus) Handle(input task_usecase.GetResumeStatusInput) (*task_usecase.OutputResume, error) {
	args := m.Called(input)
	var result *task_usecase.OutputResume
	if args.Get(0) != nil {
		result = args.Get(0).(*task_usecase.OutputResume)
	}
	return result, args.Error(1)
}

func TestTaskResponseFromOutput(t *testing.T) {
	t.Run("should return a task_usecase response from output", func(t *testing.T) {
		dateString := "2019-01-01T00:00:00"
		date, _ := time.Parse("2006-01-02T15:04:05", dateString)
		output := task_usecase.Output{
			ID:          "id",
			Title:       "title",
			Description: "description",
			DateEnd:     &date,
			DateInit:    &date,
			Status:      task.Open,
		}

		got := taskResponseFromOutput(output)

		assert.NotNil(t, got)
		assert.Equal(t, output.Title, got.Title)
	})

	t.Run("should return a empty task_usecase", func(t *testing.T) {
		got := taskResponseFromOutput(task_usecase.Output{})

		assert.Equal(t, got.Title, "")
	})
}

func TestTaskResponseFromOutputs(t *testing.T) {
	t.Run("should return a list of tasks", func(t *testing.T) {
		got := taskResponseFromOutputs([]task_usecase.Output{{}, {}})

		assert.Len(t, *got, 2)
	})
}

func TestTaskController_Create(t *testing.T) {
	date := time.Now()

	title, description, token := "title", "description", "the_token"

	output := &task_usecase.Output{
		ID:          "id",
		Title:       "title",
		Description: "description",
		DateEnd:     &date,
		DateInit:    &date,
		Status:      task.Open,
	}

	t.Run("should register a task_usecase", func(t *testing.T) {

		request := Request{
			LoggedUser: &security.User{},
			Body:       fmt.Sprintf("{\"date_init\":\"%s\",\"date_end\":\"%s\",\"title\":\"%s\",\"description\":\"%s\"}", date, date, title, description),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockTCCreate := new(mockTCCreate)
		mockTCCreate.On("Handle", mock.Anything).Return(output, nil)

		uc := TaskController{Authenticator: mockAuth, TCCreate: mockTCCreate}

		response := uc.Create(request)

		assert.NotNil(t, response)
		assert.Equal(t, response.HttpCode, httpGO.StatusCreated)
	})

	t.Run("should return a error a task_usecase when object is invalid", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\":\"%s\",\"date_init\":\"%s\",\"date_end\":\"%s\",\"title\":\"%s\",\"description\":\"%s\"}", "", date, date, title, description),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockTCCreate := new(mockTCCreate)
		mockTCCreate.On("Handle", mock.Anything).Return(nil, task.ErrIdUserIsInvalid)

		uc := TaskController{Authenticator: mockAuth, TCCreate: mockTCCreate}
		response := uc.Create(request)

		assert.NotNil(t, response)
		assert.Equal(t, response.HttpCode, httpGO.StatusInternalServerError)
	})
}

func TestTaskController_Get(t *testing.T) {
	date := time.Now()

	title, description, token := "title", "description", "the_token"

	output := &task_usecase.Output{
		ID:          "id",
		Title:       title,
		Description: description,
		DateEnd:     &date,
		DateInit:    &date,
		Status:      task.Open,
	}

	t.Run("should get a task_usecase valid", func(t *testing.T) {

		request := Request{
			Body: fmt.Sprintf("{\"id\":\"%s\"}", "123"),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockTCGet := new(mockTCGet)
		mockTCGet.On("Handle", mock.Anything).Return(output, nil)

		uc := TaskController{Authenticator: mockAuth, TCGet: mockTCGet}

		response := uc.Get(request)

		assert.NotNil(t, response)
		assert.Equal(t, response.HttpCode, httpGO.StatusOK)
	})

	t.Run("should return error when don't exists task_usecase", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"id\":\"%s\"}", "123"),
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockTCGet := new(mockTCGet)
		mockTCGet.On("Handle", mock.Anything).Return(nil, repository.ErrTaskNotFound)

		uc := TaskController{Authenticator: mockAuth, TCGet: mockTCGet}

		response := uc.Get(request)

		assert.NotNil(t, response)
		assert.Equal(t, response.HttpCode, httpGO.StatusNotFound)
	})
}

func TestTaskController_GetAllByClientId(t *testing.T) {

	t.Run("should get all taks by client", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"id_user\":\"%s\"}", "123"),
		}
		date := time.Now()

		token := "the_token"

		output := &[]task_usecase.Output{{
			ID:          "id",
			Title:       "title",
			Description: "description",
			DateEnd:     &date,
			DateInit:    &date,
			Status:      task.Open,
		},
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockTCGetAllByClientId := new(mockTCGetAllByClientId)
		mockTCGetAllByClientId.On("Handle", mock.Anything).Return(output, nil)

		uc := TaskController{Authenticator: mockAuth, TCGetAllByClient: mockTCGetAllByClientId}
		response := uc.GetAllByClientId(request)

		assert.Equal(t, response.HttpCode, httpGO.StatusOK)
	})

	t.Run("should return empty tasks", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"id_user\":\"%s\"}", "123"),
		}

		token := "the_token"

		output := &[]task_usecase.Output{{}}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockTCGetAllByClientId := new(mockTCGetAllByClientId)
		mockTCGetAllByClientId.On("Handle", mock.Anything).Return(output, nil)

		uc := TaskController{Authenticator: mockAuth, TCGetAllByClient: mockTCGetAllByClientId}
		response := uc.GetAllByClientId(request)

		assert.Equal(t, response.HttpCode, httpGO.StatusOK)
	})

}

func TestTaskController_Delete(t *testing.T) {

	token := "the_token"

	t.Run("should delete the task_usecase", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"id\":\"%s\"}", "123"),
		}

		output := &task_usecase.DeleteOutput{}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockTCDelete := new(mockTCDelete)
		mockTCDelete.On("Handle", mock.Anything).Return(output, nil)

		uc := TaskController{Authenticator: mockAuth, TCDelete: mockTCDelete}

		response := uc.Delete(request)

		assert.NotNil(t, response)
		assert.Equal(t, response.HttpCode, httpGO.StatusOK)
	})

	t.Run("should return a error when task_usecase not found", func(t *testing.T) {
		request := Request{
			Body: fmt.Sprintf("{\"id\":\"%s\"}", "123"),
		}

		//output := &task_usecase.DeleteOutput{}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockTCDelete := new(mockTCDelete)
		mockTCDelete.On("Handle", mock.Anything).Return(nil, repository.ErrTaskNotFound)

		uc := TaskController{Authenticator: mockAuth, TCDelete: mockTCDelete}

		response := uc.Delete(request)

		assert.NotNil(t, response)
		assert.Equal(t, response.HttpCode, httpGO.StatusInternalServerError)
	})
}

func TestTaskController_Edit(t *testing.T) {
	date := time.Now()

	title, token, description2 := "title", "the_token", "description2"

	output := &task_usecase.EditOutput{
		Title:       title,
		Description: description2,
		DateEnd:     &date,
		DateInit:    &date,
		Status:      task.Open,
		IDProject:   "",
	}

	t.Run("should edit a task", func(t *testing.T) {

		params := map[string]string{"id": "123"}

		request := Request{
			Params:     params,
			Body:       fmt.Sprintf("{\"date_init\":\"%s\",\"date_end\":\"%s\",\"title\":\"%s\",\"description\":\"%s\", \"IDProject\":\"\"}", date, date, title, description2),
			LoggedUser: &security.User{},
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockTCEdit := new(mockTCEdit)
		mockTCEdit.On("Handle", mock.Anything).Return(output, nil)

		uc := TaskController{Authenticator: mockAuth, TCEdit: mockTCEdit}

		response := uc.Edit(request)

		assert.Equal(t, response.HttpCode, httpGO.StatusOK)
	})
}

func TestTaskController_GetAllByDay(t *testing.T) {
	params := map[string]string{"day": "2006-01-02T15:04:05"}

	t.Run("should get all task by day", func(t *testing.T) {
		request := Request{
			Params:     params,
			LoggedUser: &security.User{ID: "123"},
		}

		date := time.Now()

		token := "the_token"

		output := &[]task_usecase.Output{{
			ID:          "id",
			Title:       "title",
			Description: "description",
			DateEnd:     &date,
			DateInit:    &date,
			Status:      task.Open,
		},
		}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockTCGetAllByDay := new(mockTCGetAllByDay)
		mockTCGetAllByDay.On("Handle", mock.Anything).Return(output, nil)

		uc := TaskController{Authenticator: mockAuth, TCGetAllByDay: mockTCGetAllByDay}
		response := uc.GetAllByDay(request)

		assert.Equal(t, httpGO.StatusOK, response.HttpCode)
	})

}

func TestTaskController_GetResumeStatus(t *testing.T) {
	t.Run("should resume all task", func(t *testing.T) {
		request := Request{
			LoggedUser: &security.User{ID: "123"},
		}

		token := "the_token"

		output := &task_usecase.OutputResume{Open: 1, Process: 2, Completed: 3}

		mockAuth := new(mockAuth)
		mockAuth.On("Generate", mock.Anything).Return(token, nil)
		mockTCGetResumeStatus := new(mockTCGetResumeStatus)
		mockTCGetResumeStatus.On("Handle", mock.Anything).Return(output, nil)

		uc := TaskController{Authenticator: mockAuth, TCGetResumeStatus: mockTCGetResumeStatus}

		response := uc.GetResumeStatus(request)

		assert.NotNil(t, response)
		assert.Equal(t, httpGO.StatusOK, response.HttpCode)

	})
}
