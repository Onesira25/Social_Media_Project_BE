package main

import (
	"Social_Media_Project_BE/config"
	user "Social_Media_Project_BE/features/user"
	user_data "Social_Media_Project_BE/features/user/data"
	user_handler "Social_Media_Project_BE/features/user/handler"
	user_service "Social_Media_Project_BE/features/user/service"
	"Social_Media_Project_BE/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)
	config.Migrate(db, &user.User{})

	userData := user_data.New(db)
	userService := user_service.NewService(userData)
	userHandler := user_handler.NewHandler(userService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes.InitRoute(e, userHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
