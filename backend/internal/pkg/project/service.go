package project

import (
	"backend/internal/api/presenters"
	"backend/internal/pkg/entities"
	"errors"
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

func (s *service) UpdateProject(ID string, project *entities.Project) (*presenters.Project, error) {
  projectResponse, err := s.repository.UpdateProject(ID, project)

  if err != nil {
    return nil, err
  }

  if projectResponse.HasBeenSoftDeleted {
    return nil, errors.New("project not found")
  }

  return projectResponse, nil
}

func (s *service) ReadProject(ID string) (*presenters.Project, error) {
  projectResponse, err := s.repository.ReadProject(ID)

  if err != nil {
    return nil, err
  }

  if projectResponse.HasBeenSoftDeleted {
    return nil, errors.New("project not found")
  }

	return projectResponse, nil
}

func (s *service) ReadProjectByOwner(ID string) ([]*presenters.Project, error) {
  projectsResponse := []*presenters.Project {}

  projects, err := s.repository.ReadProjectByOwner(ID)

  if err != nil {
    return nil, err
  }

  for _, project := range projects {
    if project.HasBeenSoftDeleted {
      continue
    }

    projectsResponse = append(projectsResponse, project)
  }

  return projectsResponse, nil
}


func (s *service) DeleteProject(ID string) error {
	return s.repository.DeleteProject(ID)
}
