package validation

import (
	"golang-mongo-api/models"
	"golang-mongo-api/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserCreateBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userCreateBody models.UserCreateBody

		// bind request data
		if err := c.Bind(&userCreateBody); err != nil {
			if err != nil {
				return utils.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := userCreateBody.Validate(); err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		c.Set("userCreateBody", userCreateBody)

		return next(c)
	}
}

func UserUpdateBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userUpdateBody models.UserUpdateBody

		// bind request data
		if err := c.Bind(&userUpdateBody); err != nil {
			if err != nil {
				return utils.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := userUpdateBody.Validate(); err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		c.Set("userUpdateBody", userUpdateBody)

		return next(c)
	}
}

// ValidateID ...
func ValidateID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id = c.Param("id")

		// validate id
		if _, err := primitive.ObjectIDFromHex(id); err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		c.Set("idString", id)

		return next(c)
	}
}
