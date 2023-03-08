package http

import (
	"encoding/json"
	httpGo "net/http"
	"time"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/app/security"
	"todo_project.com/internal/app/usecase/notes_usecase"
	"todo_project.com/internal/domain/task"
)

type (
	EditNotesResponse struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		DateInit    string `json:"date_init"`
		DateEnd     string `json:"date_end"`
		IDProject   string `json:"id_project"`
	}

	NotesListResponse struct {
		Notes []NotesResponse `json:"notes"`
	}

	NotesResponse struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		DateCreated string `json:"date_created"`
		DateUpdate  string `json:"date_update"`
	}

	NotesRequest struct {
		ID          string `json:"id"`
		IDUser      string `json:"id_user"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
)

func notesResponseFromOutput(output notes_usecase.Output) *NotesResponse {
	var dateUpdate, dateCreated string
	if output.DateCreated != nil {
		dateCreated = output.DateCreated.Format(DataFormat)
	}
	if output.DateUpdate != nil {
		dateUpdate = output.DateUpdate.Format(DataFormat)
	}

	return &NotesResponse{
		ID:          output.ID,
		DateCreated: dateCreated,
		DateUpdate:  dateUpdate,
		Title:       output.Title,
		Description: output.Description,
	}
}

func notesResponseFromOutputs(outputs []notes_usecase.Output) *[]NotesResponse {
	responses := make([]NotesResponse, 0, len(outputs))
	for _, output := range outputs {
		responses = append(responses, *notesResponseFromOutput(output))
	}
	return &responses
}

func notesListResponseFromOutput(outputs []notes_usecase.Output) *NotesListResponse {
	response := notesResponseFromOutputs(outputs)

	return &NotesListResponse{
		Notes: *response,
	}
}

type NotesController struct {
	NCCreate         notes_usecase.Create
	NCDelete         notes_usecase.Delete
	NCEdit           notes_usecase.Edit
	NCGet            notes_usecase.Get
	NCGetAllByClient notes_usecase.GetAllByClient
	NCGetAllByDay    notes_usecase.GetAllByDay
	NCGetByName      notes_usecase.GetByName
	Authenticator    security.Authenticator
}

func (c *NotesController) Create(request Request) Response {

	var notesBody NotesRequest
	err := json.Unmarshal([]byte(request.Body), &notesBody)
	if err != nil {
		return Response{
			HttpCode: httpGo.StatusInternalServerError,
			Body:     wrapError(err),
		}
	}

	input := notes_usecase.CreateInput{
		LoggedUser:  *request.LoggedUser,
		Title:       notesBody.Title,
		Description: notesBody.Description,
	}

	output, err := c.NCCreate.Handle(input)
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

func (c *NotesController) Get(request Request) Response {
	input := notes_usecase.GetInput{
		IDTask: request.Params["id"],
	}
	output, err := c.NCGet.Handle(input)

	if err != nil {
		httpStatus := httpGo.StatusInternalServerError
		switch err {
		case security.ErrUnauthorized:
			httpStatus = httpGo.StatusForbidden
		case repository.ErrNotesNotFound:
			httpStatus = httpGo.StatusNotFound
		}

		return Response{
			HttpCode: httpStatus,
			Body:     wrapError(err),
		}
	}

	return Response{
		HttpCode: httpGo.StatusOK,
		Body:     wrapBody(notesResponseFromOutput(*output)),
	}
}

func (c *NotesController) GetAllByClientId(request Request) Response {
	input := notes_usecase.GetAllByClientInput{
		IDUser: request.Params["id"],
	}

	output, err := c.NCGetAllByClient.Handle(input)

	if err != nil {
		httpStatus := httpGo.StatusInternalServerError
		switch err {
		case security.ErrUnauthorized:
			httpStatus = httpGo.StatusForbidden
		case repository.ErrNotesNotFound:
			httpStatus = httpGo.StatusNotFound
		}

		return Response{
			HttpCode: httpStatus,
			Body:     wrapError(err),
		}
	}

	return Response{
		HttpCode: httpGo.StatusOK,
		Body:     wrapBody(notesListResponseFromOutput(*output)),
	}

}

func (c *NotesController) Delete(request Request) Response {
	input := notes_usecase.DeleteInput{ID: request.Params["id"]}
	output, err := c.NCDelete.Handle(input)
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

func (c *NotesController) Edit(request Request) Response {
	var notesBody NotesRequest
	err := json.Unmarshal([]byte(request.Body), &notesBody)
	if err != nil {
		return Response{
			HttpCode: httpGo.StatusBadRequest,
			Body:     wrapError(err),
		}
	}

	input := notes_usecase.EditInput{
		ID:          request.Params["id"],
		IDUser:      request.LoggedUser.ID,
		Title:       notesBody.Title,
		Description: notesBody.Description,
		DateUpdate:  time.Now(),
	}

	output, err := c.NCEdit.Handle(input)

	if err != nil {
		return Response{
			HttpCode: httpGo.StatusInternalServerError,
			Body:     wrapError(err),
		}
	}

	return Response{
		HttpCode: httpGo.StatusOK,
		Body:     wrapBody(notesResponseFromOutput(*output)),
	}
}

func (c *NotesController) GetAllByDay(request Request) Response {
	dayTime, err := time.Parse(task.LayoutFromParseTime, request.Params["day"])
	if err != nil {
		return Response{
			HttpCode: httpGo.StatusBadRequest,
			Body:     wrapError(err),
		}
	}

	input := notes_usecase.GetAllByDayInput{
		IDUser: request.LoggedUser.ID,
		Day:    dayTime,
	}

	output, err := c.NCGetAllByDay.Handle(input)
	if err != nil {
		return Response{
			HttpCode: httpGo.StatusInternalServerError,
			Body:     wrapError(err),
		}
	}

	return Response{
		HttpCode: httpGo.StatusOK,
		Body:     wrapBody(notesListResponseFromOutput(*output)),
	}
}

func (c *NotesController) GetByName(request Request) Response {
	input := notes_usecase.GetByNameInput{
		NameTask: request.Params["name"],
		IDUser:   request.LoggedUser.ID,
	}

	response, err := c.NCGetByName.Handle(input)

	if err != nil {
		return Response{
			HttpCode: httpGo.StatusInternalServerError,
			Body:     wrapError(err),
		}
	}

	return Response{
		HttpCode: httpGo.StatusOK,
		Body:     wrapBody(notesListResponseFromOutput(*response)),
	}
}
