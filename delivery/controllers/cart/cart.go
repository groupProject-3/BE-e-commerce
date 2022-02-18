package cart

import (
	"be/delivery/middlewares"
	"be/delivery/templates"
	"be/lib/database/cart"
	"be/lib/database/product"
	"be/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error bad request for create new cart", nil))
		}

		_, err := cc.repo.FindId(user_id, newCart.Product_id)

		if err == nil {

			res1, err1 := cc.repo.GetById(newCart.Product_id, user_id)
			if err1 != nil {
				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for create new cart", nil))
			}

			if _, err := cc.repo.UpdateById(newCart.Product_id, user_id, newCart); err != nil {
				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for create new cart", nil))
			}

			res, err := cc.repoProd.GetById(int(newCart.Product_id))
			if err != nil {

				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for create new cart", nil))
			}

			if _, err := cc.repoProd.UpdateByIdAll(int(newCart.Product_id), templates.ProductRequest{Qty: (res.Qty + (int(res1.Qty) - int(newCart.Qty)))}); err != nil {
				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for create new cart", nil))
			}

			resCart, err := cc.repo.GetById(uint(newCart.Product_id), user_id)
			if err != nil {
				log.Info(err)
				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for update cart", nil))
			}

			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error bad request product already exist in cart and has beed updated ", resCart))
		} else {
			if _, err := cc.repo.Create(user_id, models.Cart{Product_id: newCart.Product_id, Qty: newCart.Qty}); err != nil {
				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for create cart", nil))
			}

			res, err := cc.repoProd.GetById(int(newCart.Product_id))
			if err != nil {

				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for create new cart", nil))
			}

			if _, err := cc.repoProd.UpdateByIdAll(int(newCart.Product_id), templates.ProductRequest{Qty: (res.Qty - int(newCart.Qty))}); err != nil {
				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for create new cart", nil))
			}

		}

		resCart, err := cc.repo.GetById(uint(newCart.Product_id), user_id)
		if err != nil {
			log.Info(err)
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for update cart", nil))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "success create new cart", resCart))

	}
}

func (cc *CartController) DeleteById() echo.HandlerFunc {
	return func(c echo.Context) error {
		prod_id, _ := strconv.Atoi(c.Param("prod_id"))
		user_id := uint(middlewares.ExtractTokenId(c))

		res, err := cc.repo.GetById(uint(prod_id), user_id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for delete cart", nil))
		}

		if _, err := cc.repo.DeleteById(uint(prod_id), user_id); err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for delete cart", nil))
		}

		resProd, errProd := cc.repoProd.GetById(int(prod_id))
		if errProd != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for delete cart", nil))
		}

		if _, err := cc.repoProd.UpdateByIdAll(int(prod_id), templates.ProductRequest{Qty: resProd.Qty + int(res.Qty)}); err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for delete cart", nil))
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
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error bad request for update cart", nil))
		}

		resCart, err := cc.repo.GetById(uint(prod_id), user_id)
		if err != nil {
			log.Info(err)
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for update cart", nil))
		}

		if _, err := cc.repo.UpdateById(uint(prod_id), user_id, newCart); err != nil {
			log.Info(err)
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for update cart", nil))
		}

		resProd, errProd := cc.repoProd.GetById(int(resCart.Product_id))
		if errProd != nil {
			log.Info(err)
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for update cart", nil))
		}

		if _, err := cc.repoProd.UpdateByIdAll(int(resCart.Product_id), templates.ProductRequest{Qty: resProd.Qty + (int(resCart.Qty) - int(newCart.Qty))}); err != nil {
			log.Info(err)
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for update cart", nil))
		}

		res, err2 := cc.repo.GetById(uint(prod_id), user_id)
		if err2 != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for update cart", nil))
		}


		priceTotal := res.Qty *uint(res.Price)




		return c.JSON(http.StatusAccepted, templates.Success(http.StatusAccepted, "success update cart", res))
	}
}

func (cc *CartController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := uint(middlewares.ExtractTokenId(c))

		res, err := cc.repo.GetAll(user_id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server for get all cart", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "success get all cart", res))

	}
}
