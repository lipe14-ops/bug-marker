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

  imageResponse, err := s.database.CreateImage(ID, imageEntity)

  if err != nil {
    return nil, err
  }

  return s.imagebase.GetPreSignedImageUrl(imageResponse)
}

func (s *service) UpdateImage(ID string, image *entities.Image) (*presenters.Image, error) {
  imageResponse, err := s.database.UpdateImage(ID, image)

  if err != nil {
    return nil, err
  }

  return s.imagebase.GetPreSignedImageUrl(imageResponse)
}

func (s *service) ReadImage(ID string) (*presenters.Image, error) {
  image, err := s.database.ReadImage(ID)

  if err != nil {
    return nil, err
  }

  return s.imagebase.GetPreSignedImageUrl(image)
}

func (s *service) ReadImagesByProject(ID string) ([]*presenters.Image, error) {
  imagesResponse := []*presenters.Image {}

  images, err := s.database.ReadImagesByProject(ID)

  if err != nil {
    return nil, err
  }
  
  for _, image := range images {
    preSignedImage, err := s.imagebase.GetPreSignedImageUrl(image)

    if err != nil {
      return nil, err
    }

    imagesResponse = append(imagesResponse, preSignedImage)
  }

  return imagesResponse, err
}

func (s *service) DeleteImage(ID string) error {
	return s.database.DeleteImage(ID)
}
