package project

import (
	"backend/internal/api/presenters"
	"backend/internal/pkg/entities"
)

type Service interface {
	CreateProject(ID string, project *entities.Project) (*presenters.Project, error)
	UpdateProject(ID string, project *entities.Project) (*presenters.Project, error)
	ReadProject(ID string) (*presenters.Project, error)
	ReadProjectByOwner(ID string) ([]*presenters.Project, error)
	DeleteProject(ID string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service {
		repository: repository,
	}
}

func (s *service) CreateProject(ID string, project *entities.Project) (*presenters.Project, error) {
	return s.repository.CreateProject(ID, project)
}

func (s *service) UpdateProject(ID string,  project*entities.Project) (*presenters.Project, error) {
	return s.repository.UpdateProject(ID, project)
}

func (s *service) ReadProject(ID string) (*presenters.Project, error) {
	return s.repository.ReadProject(ID)
}

func (s *service) ReadProjectByOwner(ID string) ([]*presenters.Project, error) {
	return s.repository.ReadProjectByOwner(ID)
}


func (s *service) DeleteProject(ID string) error {
	return s.repository.DeleteProject(ID)
}
