package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response map[string] interface{}

func generateResponse(data interface{}, message string) Response {
    return Response{
        "data": data,
        "message": message,
    }
}

// Response200 success
func Response200(c echo.Context, data interface{}, message string) error {
    if message == "" {
        message = "Thanh Cong"
    }
    return c.JSON(http.StatusBadRequest, generateResponse(data, message))
}

// Response400 b9ad request
 func Response400(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Dữ liệu không hợp lệ"
	}
	return c.JSON(http.StatusBadRequest, generateResponse(data, message))
}

// Response404 not found
func Response404(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Dữ liệu không tìm thấy"
	}
	return c.JSON(http.StatusNotFound, generateResponse(data, message))
}
