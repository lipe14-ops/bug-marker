package user

import (
	"backend/internal/api/presenters"
	"backend/internal/pkg/entities"
	"crypto/sha256"
	"errors"
)

type Service interface {
	CreateUser(user *entities.User) (*presenters.User, error)
	UpdateUser(ID string, user *entities.User) (*presenters.User, error)
	ReadUser(ID string) (*presenters.User, error)
	DeleteUser(ID string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service {
		repository: repository,
	}
}

func (s *service) CreateUser(user *entities.User) (*presenters.User, error) {
  hash := sha256.New()
  hash.Write([]byte(user.Password))

  user.Password = string(hash.Sum(nil))
	return s.repository.CreateUser(user)
}

func (s *service) UpdateUser(ID string, user *entities.User) (*presenters.User, error) {
  userResponse, err := s.repository.UpdateUser(ID, user)

  if err != nil {
    return nil, err
  }

  if userResponse.HasBeenSoftDeleted {
    return nil, errors.New("user not found.")
  }

  return userResponse, nil
}

func (s *service) ReadUser(ID string) (*presenters.User, error) {
  userResponse, err := s.repository.ReadUser(ID)

  if err != nil {
    return nil, err
  }

  if userResponse.HasBeenSoftDeleted {
    return nil, errors.New("user not found.")
  }

	return userResponse, nil
}

func (s *service) DeleteUser(ID string) error {
	return s.repository.DeleteUser(ID)
}
