package auth

import (
  "backend/internal/api/presenters"
)

type Service interface {
	ReadUserByCredentials(email, password string) (*presenters.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service {
		repository: repository,
	}
}

func (r *service) ReadUserByCredentials(email, password string) (*presenters.User, error) {
  return r.repository.ReadUserByCredentials(email, password)
}
