package service

import (
	"valorx-auth/internal/config"
	"valorx-auth/internal/repository"
)

type Service struct {
	UserService IUserService
}

type Option struct {
	Repository *repository.Repository
}

func InitiateService(cfg *config.Config, repository *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(repository.UserRepository),
	}
}
