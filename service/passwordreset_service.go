package service

import "project/repository"

type PasswordResetService interface {
	Create() error
}

type passwordResetService struct {
	repo repository.PasswordResetRepository
}

func NewPasswordResetService(repo repository.PasswordResetRepository) PasswordResetService {
	return &passwordResetService{repo: repo}
}

func (s *passwordResetService) Create() error {
	return nil
}
