package router

import (
	"sistem-krs/controllers"

	"github.com/labstack/echo/v4"
)

func Auth(e *echo.Echo) {
	e.GET("/login", controllers.Login)
	e.POST("/login", controllers.LoginPost)
	e.GET("/register", controllers.Register)
	e.POST("/register", controllers.RegisterPost)
	e.GET("/logout", controllers.Logout)
}