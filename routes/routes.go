package routes

import (
	"Social_Media_Project_BE/config"
	post "Social_Media_Project_BE/features/post"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, pc post.PostController) {
	postRoute(c, pc)
}

func postRoute(c *echo.Echo, pc post.PostController) {
	c.POST("/posts", pc.Create(), withJWTConfig())
	c.PUT("/posts/:postID", pc.Edit(), withJWTConfig())
	c.GET("/posts", pc.Posts())
	c.GET("/posts/:postID", pc.PostById())
	c.DELETE("/posts/:postID", pc.Delete(), withJWTConfig())
}

func withJWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	})
}
