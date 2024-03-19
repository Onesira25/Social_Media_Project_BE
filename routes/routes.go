package routes

import (
	"Social_Media_Project_BE/config"
	user "Social_Media_Project_BE/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, uc user.Controller) {
	config := echojwt.WithConfig(echojwt.Config{SigningKey: []byte(config.JWTSECRET)})

	userRoute(c, uc, config)
}

func userRoute(c *echo.Echo, uc user.Controller, config echo.MiddlewareFunc) {
	c.POST("/login", uc.Login())
	c.POST("/users", uc.Register())
	c.GET("/users", uc.Profile(), config)
	c.PUT("/users", uc.Update(), config)
	c.DELETE("/users", uc.Delete(), config)
}
