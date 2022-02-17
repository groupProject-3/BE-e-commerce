package auth

import (
	middewares "be/delivery/middlewares"
	"be/delivery/templates"
	"be/lib/database/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo auth.Auth
}

func New(repo auth.Auth) *AuthController {
	return &AuthController{
		repo: repo,
	}
}

func (ac *AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		Userlogin := templates.Userlogin{}

		if err := c.Bind(&Userlogin); err != nil || Userlogin.Email == "" || Userlogin.Password == "" {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "error in input file", nil))
		}
		checkedUser, err := ac.repo.Login(Userlogin)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error in call database", nil))
		}
		token, err := middewares.GenerateToken(checkedUser)

		if err != nil {
			return c.JSON(http.StatusNotAcceptable, templates.BadRequest(http.StatusNotAcceptable, "error in process token", nil))
		}

		return c.JSON(http.StatusOK, templates.Success(nil, "success login", map[string]interface{}{
			"data":  checkedUser,
			"token": token,
		}))
	}
}
