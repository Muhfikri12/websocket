package service

import (
	"project/domain"
	"project/repository"
)

type AuthService interface {
	Login(user domain.User) (string, string, bool, error)
	Register(user *domain.User) error
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(user domain.User) (string, bool, error) {
	return s.repo.Authenticate(user)
}

func (s *authService) Register(user *domain.User) error {
	return s.repo.Register(user)
}
