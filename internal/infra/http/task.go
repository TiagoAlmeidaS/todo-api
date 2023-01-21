package http

import (
	"encoding/json"
	httpGo "net/http"
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/app/security"
	"todo_project.com/internal/app/usecase/task_usecase"
	"todo_project.com/internal/domain/task"
)

type (
	GetAllByClientIdResponse struct {
		Tasks []TaskResponse `json:"tasks"`
	}

	GetAllByDayResponse struct {
		Tasks []TaskResponse `json:"tasks"`
	}

	EditResponse struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		DateInit    string `json:"date_init"`
		DateEnd     string `json:"date_end"`
		IDProject   string `json:"id_project"`
	}

	ResumeResponse struct {
		Open      int `json:"open"`
		Process   int `json:"process"`
		Completed int `json:"completed"`
	}

	TaskResponse struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		DateInit    string `json:"date_init"`
		DateEnd     string `json:"date_end"`
		IDProject   string `json:"id_project"`
	}

	TaskRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		IDUser      string `json:"id_user"`
		DateInit    string `json:"date_init"`
		DateEnd     string `json:"date_end"`
		IDProject   string `json:"id_project"`
	}
)

func resumeResponseFromOutPut(output task_usecase.OutputResume) *ResumeResponse {
	return &ResumeResponse{
		Open:      output.Open,
		Process:   output.Process,
		Completed: output.Completed,
	}
}

func editResponseFromOutPut(output task_usecase.EditOutput) *EditResponse {
	return &EditResponse{
		Title:       output.Title,
		Description: output.Description,
		Status:      string(output.Status),
		IDProject:   output.IDProject,
		DateInit:    output.DateInit.String(),
		DateEnd:     output.DateEnd.String(),
	}
}

func taskResponseFromOutput(output task_usecase.Output) *TaskResponse {
	var dateInit, dateEnd string
	if output.DateEnd != nil {
		dateInit = output.DateInit.Format(DataFormat)
	}
	if output.DateEnd != nil {
		dateEnd = output.DateEnd.Format(DataFormat)
	}

	return &TaskResponse{
		ID:          output.ID,
		DateInit:    dateInit,
		DateEnd:     dateEnd,
		Status:      string(output.Status),
		Title:       output.Title,
		Description: output.Description,
	}
}

func taskResponseFromOutputs(outputs []task_usecase.Output) *[]TaskResponse {
	responses := make([]TaskResponse, 0, len(outputs))
	for _, output := range outputs {
		responses = append(responses, *taskResponseFromOutput(output))
	}
	return &responses
}

func getAllByClientIdResponseFromOutputs(outputs []task_usecase.Output) *GetAllByClientIdResponse {
	response := taskResponseFromOutputs(outputs)

	return &GetAllByClientIdResponse{
		Tasks: *response,
	}
}

func getAllByDayResponseFromOutputs(outputs []task_usecase.Output) *GetAllByDayResponse {
	response := taskResponseFromOutputs(outputs)

	return &GetAllByDayResponse{
		Tasks: *response,
	}

}

type TaskController struct {
	TCCreate          task_usecase.Create
	TCDelete          task_usecase.Delete
	TCEdit            task_usecase.Edit
	TCGet             task_usecase.Get
	TCGetAllByClient  task_usecase.GetAllByClient
	TCGetAllByDay     task_usecase.GetAllByDay
	TCGetResumeStatus task_usecase.GetResumeStatus
	Authenticator     security.Authenticator
}

func (c *TaskController) Create(request Request) Response {

	var taskBody TaskRequest
	err := json.Unmarshal([]byte(request.Body), &taskBody)
	if err != nil {
		return Response{
			HttpCode: httpGo.StatusInternalServerError,
			Body:     wrapError(err),
		}
	}

	input := task_usecase.CreateInput{
		LoggedUser:  *request.LoggedUser,
		DateInit:    taskBody.DateInit,
		DateEnd:     taskBody.DateEnd,
		Title:       taskBody.Title,
		Description: taskBody.Description,
	}

	output, err := c.TCCreate.Handle(input)
	if err != nil {
		return Response{
			HttpCode: httpGo.StatusInternalServerError,
			Body:     wrapBody(err),
		}
	}

	return Response{
		HttpCode: httpGo.StatusCreated,
		Body:     wrapBody(output),
	}
}

func (c *TaskController) Get(request Request) Response {
	input := task_usecase.GetInput{
		IDTask: request.Params["id"],
	}
	output, err := c.TCGet.Handle(input)

	if err != nil {
		httpStatus := httpGo.StatusInternalServerError
		switch err {
		case security.ErrUnauthorized:
			httpStatus = httpGo.StatusForbidden
		case repository.ErrTaskNotFound:
			httpStatus = httpGo.StatusNotFound
		}

		return Response{
			HttpCode: httpStatus,
			Body:     wrapError(err),
		}
	}

	return Response{
		HttpCode: httpGo.StatusOK,
		Body:     wrapBody(taskResponseFromOutput(*output)),
	}
}

func (c *TaskController) GetAllByClientId(request Request) Response {
	input := task_usecase.GetAllByClientInput{
		IDUser: request.Params["id"],
	}

	output, err := c.TCGetAllByClient.Handle(input)

	if err != nil {
		httpStatus := httpGo.StatusInternalServerError
		switch err {
		case security.ErrUnauthorized:
			httpStatus = httpGo.StatusForbidden
		case repository.ErrTaskNotFound:
			httpStatus = httpGo.StatusNotFound
		}

		return Response{
			HttpCode: httpStatus,
			Body:     wrapError(err),
		}
	}

	return Response{
		HttpCode: httpGo.StatusOK,
		Body:     wrapBody(getAllByClientIdResponseFromOutputs(*output)),
	}

}

func (c *TaskController) Delete(request Request) Response {
	input := task_usecase.DeleteInput{ID: request.Params["id"]}
	output, err := c.TCDelete.Handle(input)
	if err != nil {
		return Response{
			HttpCode: httpGo.StatusInternalServerError,
			Body:     wrapError(err),
		}
	}

	return Response{
		HttpCode: httpGo.StatusOK,
		Body:     wrapBody(*output),
	}

}

func (c *TaskController) Edit(request Request) Response {
	var taskBody TaskRequest
	err := json.Unmarshal([]byte(request.Body), &taskBody)
	if err != nil {
		return Response{
			HttpCode: httpGo.StatusBadRequest,
			Body:     wrapError(err),
		}
	}

	input := task_usecase.EditInput{
		ID:          request.Params["id"],
		IDUser:      request.LoggedUser.ID,
		Title:       taskBody.Title,
		Description: taskBody.Description,
		Status:      task.Status(taskBody.Status),
		DateInit:    taskBody.DateInit,
		DateEnd:     taskBody.DateEnd,
	}

	output, err := c.TCEdit.Handle(input)

	if err != nil {
		return Response{
			HttpCode: httpGo.StatusInternalServerError,
			Body:     wrapError(err),
		}
	}

	return Response{
		HttpCode: httpGo.StatusOK,
		Body:     wrapBody(editResponseFromOutPut(*output)),
	}
}

func (c *TaskController) GetAllByDay(request Request) Response {
	dayTime, err := time.Parse(task.LayoutFromParseTime, request.Params["day"])
	if err != nil {
		return Response{
			HttpCode: httpGo.StatusBadRequest,
			Body:     wrapError(err),
		}
	}

	input := task_usecase.GetAllByDayInput{
		IDUser: request.LoggedUser.ID,
		Day:    dayTime,
	}

	if request.LoggedUser.ID == "" {
		return Response{
			HttpCode: httpGo.StatusUnauthorized,
		}
	}

	output, err := c.TCGetAllByDay.Handle(input)
	if err != nil {
		return Response{
			HttpCode: httpGo.StatusInternalServerError,
			Body:     wrapError(err),
		}
	}

	return Response{
		HttpCode: httpGo.StatusOK,
		Body:     wrapBody(getAllByDayResponseFromOutputs(*output)),
	}
}

func (c *TaskController) GetResumeStatus(request Request) Response {
	input := task_usecase.GetResumeStatusInput{
		IDUser: request.LoggedUser.ID,
	}

	response, err := c.TCGetResumeStatus.Handle(input)
	if err != nil {
		return Response{
			HttpCode: httpGo.StatusInternalServerError,
			Body:     wrapError(err),
		}
	}

	return Response{
		HttpCode: httpGo.StatusOK,
		Body:     wrapBody(resumeResponseFromOutPut(*response)),
	}

}
