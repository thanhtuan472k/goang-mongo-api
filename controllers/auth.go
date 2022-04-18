package controllers

import (
	"myapp/models"
	"myapp/service"
	"myapp/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserRegister(c echo.Context) error {
	return c.JSON(http.StatusOK, "Register User, Create User for admin")
}

func UserLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, "Login for Admin feature")
}

func AdminLogin(c echo.Context) error {
	var admin = c.Get("adminLoginBody").(models.AdminLoginBody)

	// process data
	token, err := service.AdminLogin(admin)

	// if error
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	// token
	data := map[string]interface{}{
		"token":   token,
		"isAdmin": true,
	}

	return utils.Response200(c, data, "")
}
