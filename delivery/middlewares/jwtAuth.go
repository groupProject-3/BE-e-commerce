package middewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func JwtMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("secret"),
	})
}
