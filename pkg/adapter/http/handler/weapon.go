package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/usecase"
)

type WeaponHandler struct {
	weaponUsecase       usecase.IWeaponUseCase
	usersWeaponsUsecase usecase.IUsersWeaponsUseCase
}

func NewWeaponHandler(weaponUsecase usecase.IWeaponUseCase, usersWeaponsUsecase usecase.IUsersWeaponsUseCase) *WeaponHandler {
	return &WeaponHandler{
		weaponUsecase:       weaponUsecase,
		usersWeaponsUsecase: usersWeaponsUsecase,
	}
}

func (wh *WeaponHandler) GetWeapons() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("user_id").(string)
		if userID == "" {
			return c.JSON(http.StatusUnauthorized, "you are not logged in")
		}
		weapons, err := wh.weaponUsecase.GetWeapons()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, weapons)
	}
}

func (wh *WeaponHandler) DrawGacha() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("user_id").(string)
		if userID == "" {
			return c.JSON(http.StatusUnauthorized, "you are not logged in")
		}
		weapon, err := wh.usersWeaponsUsecase.DrawGacha(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, weapon)
	}
}
