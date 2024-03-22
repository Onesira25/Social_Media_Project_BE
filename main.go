package main

import (
	"Social_Media_Project_BE/config"
	comment_data "Social_Media_Project_BE/features/comment/data"
	comment_handler "Social_Media_Project_BE/features/comment/handler"
	comment_services "Social_Media_Project_BE/features/comment/services"
	post_data "Social_Media_Project_BE/features/post/data"
	post_handler "Social_Media_Project_BE/features/post/handler"
	post_services "Social_Media_Project_BE/features/post/services"
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
	config.Migrate(db, &user_data.User{}, &post_data.Post{}, &comment_data.Comment{})

	userData := user_data.New(db)
	userService := user_service.NewService(userData)
	userHandler := user_handler.NewHandler(userService)

	postData := post_data.New(db)
	postService := post_services.PostService(postData)
	postHandler := post_handler.NewHandler(postService)

	commentData := comment_data.New(db)
	commentService := comment_services.CommentService(commentData)
	commentHandler := comment_handler.NewHandler(commentService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, userHandler, postHandler, commentHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
