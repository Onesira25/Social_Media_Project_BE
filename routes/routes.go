package routes

import (
	"Social_Media_Project_BE/config"
	"Social_Media_Project_BE/features/comment"
	post "Social_Media_Project_BE/features/post"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, pc post.PostController, cc comment.CommentController) {
	postRoute(c, pc)
	commentRoute(c, cc)
}

func postRoute(c *echo.Echo, pc post.PostController) {
	c.POST("/posts", pc.Create(), withJWTConfig())
	c.PUT("/posts/:postID", pc.Edit(), withJWTConfig())
	c.GET("/posts", pc.Posts())
	c.GET("/posts/:postID", pc.PostById())
	c.DELETE("/posts/:postID", pc.Delete(), withJWTConfig())
}

func commentRoute(c *echo.Echo, cc comment.CommentController) {
	c.POST("/comments", cc.Create(), withJWTConfig())
	c.DELETE("/comments/:commentID", cc.Delete(), withJWTConfig())
}

func withJWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	})
}
