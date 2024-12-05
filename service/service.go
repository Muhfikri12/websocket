package service

import "project/repository"

type Service struct {
	Auth          AuthService
	PasswordReset PasswordResetService
	User          UserService
}

func NewService(repo repository.Repository) Service {
	return Service{
		Auth:          NewAuthService(repo.Auth),
		PasswordReset: NewPasswordResetService(repo.PasswordReset),
		User:          NewUserService(repo.User),
	}
}
