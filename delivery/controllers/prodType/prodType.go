package prodType

import (
	"be/delivery/middlewares"
	"be/delivery/templates"
	"be/lib/database/prodType"
	"be/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProdTypeController struct {
	repo prodType.ProductType
}

func New(repo prodType.ProductType) *ProdTypeController {
	return &ProdTypeController{
		repo: repo,
	}
}

func (pc *ProdTypeController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := middlewares.ExtractTokenAdmin(c)

		if email != "admin@gmail.com" {
			return c.JSON(http.StatusUnauthorized, templates.BadRequest(http.StatusUnauthorized, "error not authorized", nil))
		}

		newProdT := templates.ProductTypeRequest{}

		if err := c.Bind(newProdT); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error in request for create new product type", nil))
		}

		res, err := pc.repo.Create(models.ProductType{Name: newProdT.Name})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "error internal server error for create new product type", nil))
		}

		return c.JSON(http.StatusAccepted, templates.Success(http.StatusAccepted, "success create new product Type", templates.ProductTypeResponse{ID: res.ID, CreatedAt: res.CreatedAt, UpdatedAt: res.UpdatedAt, Name: res.Name}))
	}
}

func (pc *ProdTypeController) UpdateById() echo.HandlerFunc {

	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		email := middlewares.ExtractTokenAdmin(c)

		if email != "admin@gmail.com" {
			return c.JSON(http.StatusUnauthorized, templates.BadRequest(http.StatusUnauthorized, "error not authorized", nil))
		}

		newProdT := templates.ProductTypeRequest{}

		if err := c.Bind(newProdT); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error in request for update product type", nil))
		}

		res, err := pc.repo.UpdateById(id, newProdT)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(http.StatusInternalServerError, "error internal server error for update product type", nil))
		}

		return c.JSON(http.StatusAccepted, templates.Success(http.StatusAccepted, "success create new product Type", templates.ProductTypeResponse{ID: res.ID, CreatedAt: res.CreatedAt, UpdatedAt: res.UpdatedAt, Name: res.Name}))

	}
}

