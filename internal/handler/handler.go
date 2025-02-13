package handler

import (
	"valorx-auth/internal/config"
	"valorx-auth/internal/service"
)

type Handler struct {
	UserHandler IUserHandler
}

func InitiateHandler(cfg *config.Config, services *service.Service) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(cfg, services.UserService),
	}
}
