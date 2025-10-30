package class

import (
	"backend/internal/api/presenters"
	"backend/internal/pkg/entities"
)

type Service interface {
	CreateClass(ID string, class *entities.Class) (*presenters.Class, error)
	UpdateClass(projectID string, project *entities.Class) (*presenters.Class, error)
	ReadClass(projectID, classID string) (*presenters.Class, error)
	ReadClassesByProject(ID string) ([]*presenters.Class, error)
	DeleteClass(projectID, classID string) error
}

type service struct {
  repository Repository
}

func NewService(repository Repository) Service {
  return &service {
    repository: repository,
  }
}

func (s *service) CreateClass(ID string, class *entities.Class) (*presenters.Class, error) {
  return s.repository.CreateClass(ID, class)
}

func (s *service) UpdateClass(projectID string, class *entities.Class) (*presenters.Class, error) {
  return s.repository.UpdateClass(projectID, class)
}

func (s *service) ReadClass(projectID, classID string) (*presenters.Class, error) {
  return s.repository.ReadClass(projectID, classID)
}

func (s *service) ReadClassesByProject(ID string) ([]*presenters.Class, error) {
	return s.repository.ReadClassesByProject(ID)
}

func (s *service) DeleteClass(projectID, classID string) error {
	return s.repository.DeleteClass(projectID, classID)
}
