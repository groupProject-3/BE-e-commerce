package paymentmethod

import (
	"be/delivery/middlewares"
	"be/delivery/templates"
	paymentmethod "be/lib/database/paymentMethod"
	"be/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type PaymentMethodController struct {
	repo paymentmethod.PaymentMethod
}

func New(repo paymentmethod.PaymentMethod) *PaymentMethodController {
	return &PaymentMethodController{
		repo: repo,
	}
}

func (pc *PaymentMethodController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		email := middlewares.ExtractTokenAdmin(c)
		if email != "admin@gmail.com" {
			return c.JSON(http.StatusUnauthorized, templates.BadRequest(http.StatusUnauthorized, "error not authorized", nil))
		}

		newPm := templates.PaymentMethodRequest{}

		if err := c.Bind(&newPm); err != nil || newPm.Name == "" {
			log.Info(err, newPm)
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error in request for create new payment method", nil))
		}

		res, err := pc.repo.Create(models.PaymentMethod{Name: newPm.Name})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "error internal server error for create new payment method", nil))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "success create new payment method", res))
	}
}

func (pc *PaymentMethodController) UpdateById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		email := middlewares.ExtractTokenAdmin(c)
		if email != "admin@gmail.com" {
			return c.JSON(http.StatusUnauthorized, templates.BadRequest(http.StatusUnauthorized, "error not authorized", nil))
		}

		newPm := templates.PaymentMethodRequest{}

		if err := c.Bind(&newPm); err != nil || newPm.Name == "" {
			log.Info(err, newPm)
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error in request for update payment method", nil))
		}

		res, err := pc.repo.UpdateById(id, newPm)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "error internal server error for update payment method", nil))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "success update payment method", res))
	}
}

func (pc *PaymentMethodController) DeleteById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		email := middlewares.ExtractTokenAdmin(c)
		if email != "admin@gmail.com" {
			return c.JSON(http.StatusUnauthorized, templates.BadRequest(http.StatusUnauthorized, "error not authorized", nil))
		}


		res, err := pc.repo.DeleteById(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "error internal server error for delete payment method", nil))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "success delete payment method", res))
	}
}

func (pc *PaymentMethodController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := middlewares.ExtractTokenAdmin(c)
		if email != "admin@gmail.com" {
			return c.JSON(http.StatusUnauthorized, templates.BadRequest(http.StatusUnauthorized, "error not authorized", nil))
		}


		res, err := pc.repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "error internal server error for get all payment method", nil))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "success get all payment method", res))
	}
}