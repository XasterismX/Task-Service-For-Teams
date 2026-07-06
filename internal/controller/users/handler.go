package users

import (
	"Task-Service-For-Teams/internal/entity"
	"log/slog"
)

type Handler struct {
	service UseCase
	logger  *slog.Logger
}
type UseCase interface {
	CreateUser(u entity.User) (entity.User, error)
	GetUsers(offset, limit int) ([]entity.User, error)
	GetOneUser(id int) (entity.User, error)
	UpdateUser(id int, user entity.User) (entity.User, error)
	DeleteUser(id int) error
}

func NewHandler(service UseCase, logger *slog.Logger) Handler {
	return Handler{service, logger}
}
