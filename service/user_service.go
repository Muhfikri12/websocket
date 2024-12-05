package service

import "project/repository"

type UserService interface {
	Get() error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Get() error {
	return nil
}
