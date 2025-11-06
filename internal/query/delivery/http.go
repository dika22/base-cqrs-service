package handler

import (
	"cqrs-base/internal/query/service"

	"github.com/labstack/echo/v4"
)

type UserQueryHandler struct {
	service *service.UserQueryService
	userAPI *echo.Group
}

func (h *UserQueryHandler) GetUserByID(e echo.Context) error {
	return  nil
}

func (h *UserQueryHandler) GetAllUsers(e echo.Context) error {
	return nil
}


func NewUserQueryHandler(r *echo.Group, s *service.UserQueryService) {
	u := UserQueryHandler{userAPI: r}
	r.POST("/:id", u.GetUserByID).Name = "users.signup"
	r.POST("", u.GetAllUsers).Name = "users.login"
}