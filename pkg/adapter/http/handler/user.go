package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/adapter/schemas"
	"github.com/yamato0211/plesio-server/pkg/usecase"
)

type userHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) *userHandler {
	return &userHandler{
		usecase: uu,
	}
}

func (uh *userHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		user, err := uh.usecase.GetUser(c, id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, schemas.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
}
