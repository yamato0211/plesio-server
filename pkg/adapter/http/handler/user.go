package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/adapter/schemas"
	"github.com/yamato0211/plesio-server/pkg/usecase"
)

type UserHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		usecase: uu,
	}
}

func (uh *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		user, err := uh.usecase.GetUser(c, id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, user)
	}
}
func (uh *UserHandler) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req schemas.LoginRequest
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		_, err = uh.usecase.LoginUser(c, req.Name, req.Email, req.GitID)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, "create user ok!!")
	}
}

func (uh *UserHandler) GetMyCoin() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("user_id") == nil {
			return c.JSON(http.StatusUnauthorized, "you are not logged in")
		}
		userID := c.Get("user_id").(string)
		user, err := uh.usecase.GetUser(c, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		res := &schemas.MyCoinResponse{
			Coin: user.Coin,
		}
		return c.JSON(http.StatusOK, res)
	}
}

// func (uh *UserHandler) LoginBonus() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var req schemas.LoginBonusRequest
// 		err := c.Bind(&req)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, err)
// 		}
// 		user, err := uh.usecase.LoginBonus(c, req.ID, req.GitID)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, err)
// 		}
// 		return c.JSON(http.StatusOK, user)
// 	}
// }
