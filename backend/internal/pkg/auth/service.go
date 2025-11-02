package auth

import (
  "backend/internal/api/presenters"
  "crypto/sha256"
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
  hash := sha256.New()
  hash.Write([]byte(password))
  return r.repository.ReadUserByCredentials(email, string(hash.Sum(nil)))
}
