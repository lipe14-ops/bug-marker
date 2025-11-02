package user

import (
	"backend/internal/pkg/entities"
	"backend/internal/api/presenters"
  "crypto/sha256"
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
	return s.repository.UpdateUser(ID, user)
}

func (s *service) ReadUser(ID string) (*presenters.User, error) {
	return s.repository.ReadUser(ID)
}

func (s *service) DeleteUser(ID string) error {
	return s.repository.DeleteUser(ID)
}
