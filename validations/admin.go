package validation

import (
	"golang-mongo-api/models"
	"golang-mongo-api/utils"

	"github.com/labstack/echo/v4"
)

// AdminLoginBody ...
func AdminLoginBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var adminLoginBody models.AdminLoginBody

		// bind request data
		if err := c.Bind(&adminLoginBody); err != nil {
			if err != nil {
				return utils.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := adminLoginBody.Validate(); err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		c.Set("adminLoginBody", adminLoginBody)

		return next(c)
	}
}
