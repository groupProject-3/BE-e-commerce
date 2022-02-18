package product

import (
	"be/delivery/middlewares"
	"be/delivery/templates"
	"be/lib/database/product"
	"be/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ProdController struct {
	repo product.Product
}

func New(repo product.Product) *ProdController {
	return &ProdController{
		repo: repo,
	}
}

func (pc *ProdController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := uint(middlewares.ExtractTokenId(c))

		newProd := templates.ProductRequest{}
		if err := c.Bind(&newProd); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error bad request for create new product", nil))
		}

		res, err := pc.repo.Create(user_id, models.Product{Name: newProd.Name, Product_type_id: newProd.Product_type_id, Price: newProd.Price, Qty: newProd.Qty, Description: newProd.Description})
		log.Info(err)
		log.Info(newProd)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for create product", nil))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "success create new product", templates.ProductResponse{ID: res.ID, CreatedAt: res.CreatedAt, UpdatedAt: res.UpdatedAt, Name: res.Name, Image: res.Image, Price: res.Price, Qty: res.Qty, Description: res.Description}))
	}
}

func (pc *ProdController) UpdateById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user_id := uint(middlewares.ExtractTokenId(c))

		newProd := templates.ProductRequest{}
		if err := c.Bind(&newProd); err != nil {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error bad request for update product", nil))
		}

		res, err := pc.repo.UpdateById(id, user_id, newProd)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for update product", nil))
		}

		return c.JSON(http.StatusAccepted, templates.Success(http.StatusAccepted, "success update product", templates.ProductResponse{ID: res.ID, CreatedAt: res.CreatedAt, UpdatedAt: res.UpdatedAt, Name: res.Name, Image: res.Image, Price: res.Price, Qty: res.Qty, Description: res.Description}))
	}
}

func (pc *ProdController) DeleteById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user_id := uint(middlewares.ExtractTokenId(c))

		res, err := pc.repo.DeleteById(id, user_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for delete product", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "success delete product", res))
	}
}

func (pc *ProdController) GetAllMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := uint(middlewares.ExtractTokenId(c))

		res, err := pc.repo.GetAllMe(user_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for get all my product", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "success get all my product", res))
	}
}

func (pc *ProdController) GetByIdMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user_id := uint(middlewares.ExtractTokenId(c))

		res, err := pc.repo.GetByIdMe(id, user_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for get my product", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "success get my product", res))
	}
}

func (pc *ProdController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := pc.repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for get all product", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "success get all product", res))
	}
}

func (pc *ProdController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		res, err := pc.repo.GetById(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for get product", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "success get product", res))
	}
}
