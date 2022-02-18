package main

import (
	"be/configs"
	"be/delivery/controllers/auth"
	"be/delivery/controllers/cart"
	"be/delivery/controllers/order"
	"be/delivery/controllers/paymentmethod"
	"be/delivery/controllers/prodType"
	"be/delivery/controllers/product"
	"be/delivery/controllers/user"
	"be/delivery/routes"
	authlib "be/lib/database/auth"
	cartLib "be/lib/database/cart"
	orderLib "be/lib/database/order"
	paymentmethodLib "be/lib/database/paymentmethod"
	prodTypeLib "be/lib/database/prodType"
	prodLib "be/lib/database/product"
	userlib "be/lib/database/user"
	"be/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := userlib.New(db)
	userController := user.New(userRepo)

	prodTypeRepo := prodTypeLib.New(db)
	prodTypeController := prodType.New(prodTypeRepo)

	prodRepo := prodLib.New(db)
	prodController := product.New(prodRepo)

	cartRepo := cartLib.New(db)
	cartController := cart.New(cartRepo, prodRepo)

	pmRepo := paymentmethodLib.New(db)
	pmController := paymentmethod.New(pmRepo)

	orderRepo := orderLib.New(db)
	orderController := order.New(orderRepo)

	authRepo := authlib.New(db)
	authController := auth.New(authRepo)

	e := echo.New()

	routes.UserPath(e, userController, authController)
	routes.AdminPath(e, userController)
	routes.ProductTypePath(e, prodTypeController)
	routes.ProductPath(e, prodController)
	routes.CartPath(e, cartController)
	routes.PaymentMethodPath(e, pmController)
	routes.OrderPath(e, orderController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
