package order

import (
	"be/delivery/middlewares"
	"be/delivery/templates"
	"be/lib/database/order"
	"be/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	repo order.Order
}

func New(repo order.Order) *OrderController {
	return &OrderController{
		repo: repo,
	}
}

func (oc *OrderController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := uint(middlewares.ExtractTokenId(c))

		newOr := templates.OrderRequest{}

		if err := c.Bind(&newOr); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error bad request for create new order", nil))
		}

		res, err := oc.repo.Create(user_id, models.Order{Payment_method_id: newOr.Payment_method_id})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for create new order", nil))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "success create new order", res))
	}
}

func (oc *OrderController) DeleteById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user_id := uint(middlewares.ExtractTokenId(c))

		res, err := oc.repo.DeleteById(uint(id), user_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for delete order", nil))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "success delete order", res))
	}
}

func (oc *OrderController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user_id := uint(middlewares.ExtractTokenId(c))

		res, err := oc.repo.GetById(uint(id), user_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for get order", nil))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "success get order", res))
	}
}
