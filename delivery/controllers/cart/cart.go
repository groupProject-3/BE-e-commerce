package cart

import (
	"be/delivery/middlewares"
	"be/delivery/templates"
	"be/lib/database/cart"
	"be/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type CartController struct {
	repo cart.Cart
}

func New(repo cart.Cart) *CartController {
	return &CartController{
		repo: repo,
	}
}

func (cc *CartController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := uint(middlewares.ExtractTokenId(c))

		newCart := templates.CartRequest{}
		if err := c.Bind(&newCart); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error bad request for create new cart", err))
		}

		res, err := cc.repo.CreateNew(user_id, models.Cart{Product_id: newCart.Product_id, Qty: newCart.Qty, Status: newCart.Status})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for add cart", err))

		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "success create new cart", res))

	}
}

func (cc *CartController) DeleteById() echo.HandlerFunc {
	return func(c echo.Context) error {
		prod_id, _ := strconv.Atoi(c.Param("prod_id"))
		user_id := uint(middlewares.ExtractTokenId(c))

		res, err := cc.repo.DeleteNew(uint(prod_id), user_id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for delete cart", err))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "success delete cart", res))

	}
}

func (cc *CartController) UpdateById() echo.HandlerFunc {
	return func(c echo.Context) error {
		prod_id, _ := strconv.Atoi(c.Param("prod_id"))
		user_id := uint(middlewares.ExtractTokenId(c))

		newCart := templates.CartRequest{}
		if err := c.Bind(&newCart); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error bad request for update cart", err))
		}

		res, err := cc.repo.UpdateCart(uint(prod_id), user_id, newCart)
		if err != nil {
			log.Info(err)
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for update cart", err))
		}

		return c.JSON(http.StatusAccepted, templates.Success(http.StatusAccepted, "success update cart", res))
	}
}

func (cc *CartController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		// log.Info(c.Param("status"))
		status := c.Param("filter")
		user_id := uint(middlewares.ExtractTokenId(c))
		// log.Info(status)
		// statusNow := "cart"
		res, err := cc.repo.GetAll(user_id, status)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for get all cart", err))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "success get all cart", res))

	}
}
