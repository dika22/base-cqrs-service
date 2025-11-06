package handler

import (
	"cqrs-base/internal/command/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *service.UserService
}


func (h *UserHandler) CreateUser(e echo.Context) error {
	return nil
}

func (h *UserHandler) UpdateUser(e echo.Context) error {
	return nil
}

func (h *UserHandler) DeleteUser(e echo.Context) error {
	return nil
}

func NewUserHandler(userAPI *echo.Group, s *service.UserService)  {
	u := UserHandler{service: s}
	userAPI.POST("", u.CreateUser).Name = "users.create"
	userAPI.PUT("/:id", u.UpdateUser).Name = "users.update"
	userAPI.DELETE("/:id", u.DeleteUser).Name = "users.delete"
}