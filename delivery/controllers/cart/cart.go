package cart

import (
	"be/delivery/middlewares"
	"be/delivery/templates"
	"be/lib/database/cart"
	"be/lib/database/product"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartController struct {
	repo     cart.Cart
	repoProd product.Product
}

func New(repo cart.Cart, repoProd product.Product) *CartController {
	return &CartController{
		repo:     repo,
		repoProd: repoProd,
	}
}

func (cc *CartController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := uint(middlewares.ExtractTokenId(c))

		newCart := templates.CartRequest{}
		if err := c.Bind(&newCart); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error bad request for create new product", nil))
		}

		res, err := cc.repo.CheckProduct(user_id, newCart.Product_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error in checked product for create new cart", nil))
		}

		if res == true {
			res, err := cc.repo.FindId(user_id, newCart.Product_id)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for create new cart", nil))
			}

			if _, err := cc.repo.UpdateById(res, user_id, templates.CartRequest{newCart.Product_id, newCart.Qty, newCart.Status}); err != nil {
				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for create new cart", nil))
			}

			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error bad request product already exist in cart and has beed updated ", nil))

		}

		res, 

	}
}
