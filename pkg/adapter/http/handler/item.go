package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/adapter/schemas"
	"github.com/yamato0211/plesio-server/pkg/usecase"
)

type ItemHandler struct {
	itemUsecase       usecase.IItemUsecase
	usersItemsUsecase usecase.IUsersItemsUseCase
}

func NewItemHandler(iu usecase.IItemUsecase, uiu usecase.IUsersItemsUseCase) *ItemHandler {
	return &ItemHandler{
		itemUsecase:       iu,
		usersItemsUsecase: uiu,
	}
}

func (ih *ItemHandler) GetAllItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("user_id") == nil {
			return c.JSON(http.StatusUnauthorized, "you are not logged in")
		}
		items, err := ih.itemUsecase.GetAllItem()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		res := make([]*schemas.Item, len(items))
		for i, item := range items {
			res[i] = &schemas.Item{
				ID:          item.ID,
				Name:        item.Name,
				Type:        item.Type,
				Amount:      item.Amount,
				Description: item.Description,
				Reality:     item.Reality,
				Price:       item.Price,
				CreatedAt:   item.CreatedAt,
				UpdatedAt:   item.UpdatedAt,
			}
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (ih *ItemHandler) GetAllMyItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("user_id") == nil {
			return c.JSON(http.StatusUnauthorized, "you are not logged in")
		}
		userID := c.Get("user_id").(string)
		items, err := ih.itemUsecase.GetAllMyItem(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		res := make([]*schemas.MyItemsResponce, len(items))
		for i, item := range items {
			res[i] = &schemas.MyItemsResponce{
				ID:          item.ID,
				Name:        item.Name,
				Type:        item.Type,
				Amount:      item.Amount,
				Description: item.Description,
				Reality:     item.Reality,
				Price:       item.Price,
				Count:       item.Count,
				CreatedAt:   item.CreatedAt,
				UpdatedAt:   item.UpdatedAt,
			}
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (ih *ItemHandler) BuyItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("user_id") == nil {
			return c.JSON(http.StatusUnauthorized, "you are not logged in")
		}
		userID := c.Get("user_id").(string)
		req := &schemas.BuyItemRequest{}
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		err := ih.usersItemsUsecase.BuyItem(userID, req.ItemID, req.Count)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, "success")
	}
}
