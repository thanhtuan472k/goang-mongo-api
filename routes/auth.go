package routes

import (
	"github.com/labstack/echo/v4"
	"myapp/controllers"
	"myapp/validations"
)

func Auth(e *echo.Echo) {
	auth := e.Group("/auth")
	//auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)
	auth.POST("/admin-login", controllers.AdminLogin, validation.AdminLoginBody)
}
