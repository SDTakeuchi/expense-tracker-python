package handler

import (
	"net/http"
	"strconv"

	"app/usecase"

	"github.com/labstack/echo"
	null "gopkg.in/guregu/null.v4"
)

// ItemHandler item handlerのinterface
type ItemHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type itemHandler struct {
	itemUsecase usecase.ItemUsecase
}

// NewItemHandler item handlerのコンストラクタ
func NewItemHandler(itemUsecase usecase.ItemUsecase) ItemHandler {
	return &itemHandler{itemUsecase: itemUsecase}
}

type requestItem struct {
	Name            null.String `json:"name"`
	Memo            null.String `json:"memo"`
	Price           null.Int    `json:"price"`
	MinPrice        null.Int    `json:"min_price"`
	MaxPrice        null.Int    `json:"max_price"`
	PurchaseDate    null.Time   `json:"purchase_date"`
	MinPurchaseDate null.Time   `json:"min_purchase_date"`
	MaxPurchaseDate null.Time   `json:"max_purchase_date"`
	SmallCategoryId null.String `json:"small_category_id"`
	LargeCategoryId null.String `json:"large_category_id"`
}

type responseItem struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Post itemを保存するときのハンドラー
func (th *itemHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestItem
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdItem, err := th.itemUsecase.Create(
			req.Name,
			req.Memo,
            req.Price,
            req.PurchaseDate,
            req.SmallCategoryId,
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseItem{
			ID:      createdItem.ID,
			Title:   createdItem.Title,
			Content: createdItem.Content,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get itemを取得するときのハンドラー
func (th *itemHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundItem, err := th.itemUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseItem{
			ID:      foundItem.ID,
			Title:   foundItem.Title,
			Content: foundItem.Content,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Put itemを更新するときのハンドラー
func (th *itemHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestItem
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedItem, err := th.itemUsecase.Update(id, req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseItem{
			ID:      updatedItem.ID,
			Title:   updatedItem.Title,
			Content: updatedItem.Content,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete itemを削除するときのハンドラー
func (th *itemHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = th.itemUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
