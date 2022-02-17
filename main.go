package main

import (
	"be/configs"
	"be/delivery/controllers/auth"
	"be/delivery/controllers/user"
	"be/delivery/routes"
	authlib "be/lib/database/auth"
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

	authRepo := authlib.New(db)
	AuthController := auth.New(authRepo)

	e := echo.New()

	routes.UserPath(e, userController, AuthController)
	routes.AdminPath(e, userController, AuthController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
