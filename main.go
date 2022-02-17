package main

import (
	"be/configs"
	"be/delivery/controllers/auth"
	"be/delivery/controllers/prodType"
	"be/delivery/controllers/product"
	"be/delivery/controllers/user"
	"be/delivery/routes"
	authlib "be/lib/database/auth"
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

	authRepo := authlib.New(db)
	authController := auth.New(authRepo)

	e := echo.New()

	routes.UserPath(e, userController, authController)
	routes.AdminPath(e, userController)
	routes.ProductTypePath(e, prodTypeController)
	routes.ProductPath(e, prodController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
