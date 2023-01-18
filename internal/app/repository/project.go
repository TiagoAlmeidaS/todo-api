package repository

import (
	"errors"
	"todo_project.com/internal/domain/project"
)

var (
	ErrProjectNotFound = errors.New("project not found")
)

type IProjectRepository interface {
	GetById(id string) (*project.Project, error)
	GetAllByClientId(clientId string) (*[]project.Project, error)
	Inset(project project.Project) (*project.Project, error)
	Update(project project.Project) (*project.Project, error)
	DeleteById(id string) error
}
