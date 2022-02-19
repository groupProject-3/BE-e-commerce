package user

import (
	"be/delivery/middlewares"
	"be/delivery/templates"
	"be/lib/database/user"
	"be/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo user.User
}

func New(repo user.User) *UserController {
	return &UserController{
		repo: repo,
	}
}

func (uc *UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		newUser := templates.UserRequest{}
		if err := c.Bind(&newUser); err != nil || newUser.Email == "" || newUser.Password == "" {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error in request for create new user", err))
		}
		res, err := uc.repo.Create(models.User{Name: newUser.Name, Email: newUser.Email, Password: newUser.Password})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error fo create new user", err))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "Success create new user", templates.UserResponse{ID: res.ID, CreatedAt: res.CreatedAt, UpdatedAt: res.UpdatedAt, Name: res.Name, Email: res.Email}))
	}
}

func (uc *UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := uc.repo.GetAll()
		email := middlewares.ExtractTokenAdmin(c)

		if err != nil || email != "admin@gmail.com" {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error for get all user", err))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "Success get all user", res))

	}
}
