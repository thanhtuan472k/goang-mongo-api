package routes

import (
	"myapp/controllers"
	validation "myapp/validations"

	"github.com/labstack/echo/v4"
)

func Auth(e *echo.Echo) {
	auth := e.Group("/auth")
	//auth.POST("/register", controllers.Register)
	// auth.POST("/login", controllers.UserLogin)
	auth.POST("/admin-login", controllers.AdminLogin, validation.AdminLoginBody)
}
