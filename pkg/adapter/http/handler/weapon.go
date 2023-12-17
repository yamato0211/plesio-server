package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/adapter/schemas"
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
		log.Println(weapons)
		log.Println("err: ", err)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, weapons)
	}
}

func (wh *WeaponHandler) GetAllMyWeapon() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("user_id").(string)
		if userID == "" {
			return c.JSON(http.StatusUnauthorized, "you are not logged in")
		}
		weapons, err := wh.weaponUsecase.GetWeaponByID(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		res := make([]*schemas.MyWeaponsResponce, len(weapons))
		for i, weapon := range weapons {
			res[i] = &schemas.MyWeaponsResponce{
				ID:          weapon.ID,
				Name:        weapon.Name,
				Type:        weapon.Type,
				Description: weapon.Description,
				Reality:     weapon.Reality,
				Atk:         weapon.Atk,
				CreatedAt:   weapon.CreatedAt,
				UpdatedAt:   weapon.UpdatedAt,
			}
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
		res := &schemas.Weapon{
			ID:          weapon.ID,
			Name:        weapon.Name,
			Type:        weapon.Type,
			Description: weapon.Description,
			Reality:     weapon.Reality,
			Atk:         weapon.Atk,
			CreatedAt:   weapon.CreatedAt,
			UpdatedAt:   weapon.UpdatedAt,
		}
		return c.JSON(http.StatusOK, res)
	}
}
