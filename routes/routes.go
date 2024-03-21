package routes

import (
	"Social_Media_Project_BE/config"
	comment "Social_Media_Project_BE/features/comment"
	post "Social_Media_Project_BE/features/post"
	user "Social_Media_Project_BE/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, uc user.Controller, pc post.PostController, cc comment.CommentController) {
	config := echojwt.WithConfig(echojwt.Config{SigningKey: []byte(config.JWTSECRET)})

	userRoute(c, uc, config)
	postRoute(c, pc, config)
	commentRoute(c, cc, config)
}

func userRoute(c *echo.Echo, uc user.Controller, config echo.MiddlewareFunc) {
	c.POST("/login", uc.Login())
	c.POST("/users", uc.Register())
	c.GET("/users", uc.Profile(), config)
	c.PUT("/users", uc.Update(), config)
	c.DELETE("/users", uc.Delete(), config)
}

func postRoute(c *echo.Echo, pc post.PostController, config echo.MiddlewareFunc) {
	c.POST("/posts", pc.Create(), config)
	c.PUT("/posts/:postID", pc.Edit(), config)
	c.GET("/posts", pc.Posts())
	c.GET("/posts/:postID", pc.PostById())
	c.DELETE("/posts/:postID", pc.Delete(), config)
}

func commentRoute(c *echo.Echo, cc comment.CommentController, config echo.MiddlewareFunc) {
	c.POST("/comments", cc.Create(), config)
	c.DELETE("/comments/:commentID", cc.Delete(), config)
}
