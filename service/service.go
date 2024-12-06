package service

import "project/repository"

type Service struct {
	Auth   AuthService
	Banner ServiceBanner
}

func NewService(repo repository.Repository) Service {
	return Service{
		Auth:   NewAuthService(repo.Auth),
		Banner: NewServiceBanner(repo.Banner),
	}
}
