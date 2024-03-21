package main

import (
	"Social_Media_Project_BE/config"
	comment_data "Social_Media_Project_BE/features/comment/data"
	comment_handler "Social_Media_Project_BE/features/comment/handler"
	comment_services "Social_Media_Project_BE/features/comment/services"
	post_data "Social_Media_Project_BE/features/post/data"
	post_handler "Social_Media_Project_BE/features/post/handler"
	post_services "Social_Media_Project_BE/features/post/services"
	"Social_Media_Project_BE/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	postData := post_data.New(db)
	postService := post_services.NewTodoService(postData)
	postHandler := post_handler.NewHandler(postService)

	commentData := comment_data.New(db)
	commentService := comment_services.NewTodoService(commentData)
	commentHandler := comment_handler.NewHandler(commentService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, postHandler, commentHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
