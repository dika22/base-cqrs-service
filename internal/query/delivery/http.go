package handler

import (
	"cqrs-base/internal/query/service"
	"cqrs-base/package/response"

	"github.com/labstack/echo/v4"
)

type UserQueryHandler struct {
	service *service.UserQueryService
}

func (h *UserQueryHandler) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	res , err := h.service.GetUserByID(ctx, c.Param("id")); if err != nil {
		return err
	}
	return response.JSONSuccess(c, res, "success")
}

func (h *UserQueryHandler) GetAllUsers(e echo.Context) error {
	ctx := e.Request().Context()
	res, err := h.service.GetAllUsers(ctx); if err != nil {
		return err
	}
	return response.JSONSuccess(e, res, "success")
}


func NewUserQueryHandler(r *echo.Group, s *service.UserQueryService) {
	u := UserQueryHandler{service: s}
	r.POST("/:id", u.GetUserByID).Name = "users.get-user-by-id"
	r.POST("", u.GetAllUsers).Name = "users.get-all-users"
}