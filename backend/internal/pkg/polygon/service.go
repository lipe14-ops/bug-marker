package polygon

import (
	"backend/internal/api/presenters"
	"backend/internal/pkg/entities"
)

type Service interface {
	CreatePolygon(imageID, classID string, class *entities.Polygon) (*presenters.Polygon, error)
	UpdatePolygon(imageID, polygonID, classID string, project *entities.Polygon) (*presenters.Polygon, error)
	ReadPolygon(projectID, classID string) (*presenters.Polygon, error)
	ReadPolygonsByImage(ID string) ([]*presenters.Polygon, error)
	DeletePolygon(projectID, classID string) error
}

type service struct {
  repository Repository
}

func NewService(repository Repository) Service {
  return &service {
    repository: repository,
  }
}

func (s *service) CreatePolygon(imageID, classID string, polygon *entities.Polygon) (*presenters.Polygon, error) {
  return s.repository.CreatePolygon(imageID, classID, polygon)
}

func (s *service) UpdatePolygon(imageID, polygonID, classID string, polygon *entities.Polygon) (*presenters.Polygon, error) {
  return s.repository.UpdatePolygon(imageID, polygonID, classID, polygon)
}

func (s *service) ReadPolygon(imageID, polygonID string) (*presenters.Polygon, error) {
  return s.repository.ReadPolygon(imageID, polygonID)
}

func (s *service) ReadPolygonsByImage(ID string) ([]*presenters.Polygon, error) {
	return s.repository.ReadPolygonsByImage(ID)
}

func (s *service) DeletePolygon(imageID, polygonID string) error {
	return s.repository.DeletePolygon(imageID, polygonID)
}

