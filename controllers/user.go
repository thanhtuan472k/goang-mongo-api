package controllers

import (
	"github.com/labstack/echo/v4"
	"golang-mongo-api/models"
	"golang-mongo-api/service"
	"golang-mongo-api/utils"
)

func CreateUser(c echo.Context) error {
	var userCreateBody = c.Get("userCreateBody").(models.UserCreateBody)

	// process data
	err := service.CreateUser(userCreateBody)

	// if error
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	// success
	return utils.Response200(c, nil, "")
}

func UpdateUser(c echo.Context) error {
	var (
		idString       = c.Get("idString").(string)
		userUpdateBody = c.Get("userUpdateBody").(models.UserUpdateBody)
	)

	if err := service.UpdateUser(idString, userUpdateBody); err != nil {
		return utils.Response400(c, nil, err.Error())
	}
	return utils.Response200(c, nil, "")
}

func DeleteUser(c echo.Context) error {
	var idString = c.Get("idString").(string)

	if err := service.DeleteUser(idString); err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, nil, "")
}

func GetListUser(c echo.Context) error {
	users := service.GetListUser()

	return utils.Response200(c, users, "")
}

//func GetListUserPerPage(c echo.Context) error {
//	page, _ := strconv.Atoi(c.QueryParam("page"))
//	limit, _ := strconv.Atoi(c.QueryParam("limit"))
//
//	players := service.GetListUserPerPage(page, limit)
//
//	return utils.Response200(c, players, "")
//}
