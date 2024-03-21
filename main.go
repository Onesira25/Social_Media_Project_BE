package main

import (
	"Social_Media_Project_BE/config"
	pd "Social_Media_Project_BE/features/post/data"
	ph "Social_Media_Project_BE/features/post/handler"
	ps "Social_Media_Project_BE/features/post/services"
	"Social_Media_Project_BE/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	postData := pd.New(db)
	postService := ps.NewTodoService(postData)
	postHandler := ph.NewHandler(postService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, postHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
