package image

import (
	"mime/multipart"

	"backend/internal/api/presenters"
	"backend/internal/pkg/entities"

	"backend/internal/pkg/image/repositories"
)

type Service interface {
	CreateImage(ID string, image *multipart.FileHeader) (*presenters.Image, error)
	UpdateImage(ID string, image *entities.Image) (*presenters.Image, error)
	ReadImage(ID string) (*presenters.Image, error)
	ReadImagesByProject(ID string) ([]*presenters.Image, error)
	DeleteImage(ID string) error
}

type service struct {
	imagebase repositories.ImageRepository
	database  repositories.Repository
}

func NewService(imagebase repositories.ImageRepository, database repositories.Repository) Service {
	return &service {
		imagebase: imagebase,
		database: database,
	}
}

func (s *service) CreateImage(ID string, image *multipart.FileHeader) (*presenters.Image, error) {
	imageEntity, err :=	s.imagebase.UploadImage(image)

	if err != nil {
		return nil, err
	}

	return s.database.CreateImage(ID, imageEntity)
}

func (s *service) UpdateImage(ID string, image *entities.Image) (*presenters.Image, error) {
	return s.database.UpdateImage(ID, image)
}

func (s *service) ReadImage(ID string) (*presenters.Image, error) {
	return s.database.ReadImage(ID)
}

func (s *service) ReadImagesByProject(ID string) ([]*presenters.Image, error) {
	return s.database.ReadImagesByProject(ID)
}

func (s *service) DeleteImage(ID string) error {
	return s.database.DeleteImage(ID)
}
