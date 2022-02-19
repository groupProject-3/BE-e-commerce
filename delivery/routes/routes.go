package routes

import (
	"be/delivery/controllers/auth"
	"be/delivery/controllers/cart"
	"be/delivery/controllers/order"
	paymentmethod "be/delivery/controllers/paymethod"
	"be/delivery/controllers/prodType"
	"be/delivery/controllers/product"
	"be/delivery/controllers/user"
	"be/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	e.POST("/user", uc.Create())
	e.POST("/login", ac.Login())
}

func ProductTypePath(e *echo.Echo, pc *prodType.ProdTypeController) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	e.POST("/product/type", pc.Create(), middlewares.JwtMiddleware())
	e.PUT("/product/type/:id", pc.Put(), middlewares.JwtMiddleware())
	e.DELETE("/product/type/:id", pc.Delete(), middlewares.JwtMiddleware())
	e.GET("/product/type", pc.GetAll(), middlewares.JwtMiddleware())
}

func ProductPath(e *echo.Echo, pc *product.ProdController) {

	g := e.Group("/product", middlewares.JwtMiddleware(), middleware.RemoveTrailingSlash())

	g.Use(middleware.CORS())
	// g.Pre(middleware.RemoveTrailingSlash())
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))
	g.POST("/me", pc.Create())
	g.PUT("/me/:id", pc.UpdateById())
	g.DELETE("/me/:id", pc.DeleteById())
	g.GET("/me", pc.GetAllMe())
	g.GET("/me/:id", pc.GetByIdMe())

	f := e.Group("/product", middleware.RemoveTrailingSlash())

	f.Use(middleware.CORS())
	f.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	f.GET("/all", pc.GetAll())
	f.GET("/all/:id", pc.GetById())

}

func CartPath(e *echo.Echo, cc *cart.CartController) {
	g := e.Group("/cart/me", middlewares.JwtMiddleware(), middleware.RemoveTrailingSlash())

	g.Use(middleware.CORS())
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	g.POST("", cc.Create())
	g.DELETE("/:prod_id", cc.DeleteById())
	g.PUT("/:prod_id", cc.UpdateById())
	g.GET("", cc.GetAll())
}

func PaymentMethodPath(e *echo.Echo, pc *paymentmethod.PaymentMethodController) {
	g := e.Group("/payment/method", middlewares.JwtMiddleware(), middleware.RemoveTrailingSlash())

	g.Use(middleware.CORS())
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	g.POST("", pc.Create())
	g.DELETE("/:id", pc.DeleteById())
	g.PUT("/:id", pc.UpdateById())
	g.GET("/:id", pc.GetAll())
}

func OrderPath(e *echo.Echo, oc *order.OrderController)  {
	g := e.Group("/order/me", middlewares.JwtMiddleware(), middleware.RemoveTrailingSlash())

	g.Use(middleware.CORS())
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	g.POST("", oc.Create())
	g.DELETE("", oc.DeleteById())
	g.GET("/:id", oc.GetAll())
}

func AdminPath(e *echo.Echo, uc *user.UserController) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	e.GET("/admin/user", uc.GetAll(), middlewares.JwtMiddleware())
}
