package httputils

import "github.com/labstack/echo/v4"

type (
	HttpErrorResponse struct {
		Message string `json:"message"`
	}
)

func NewHttpErrorResponse(c echo.Context, httpStatus int, message string) error {
	error := HttpErrorResponse{Message: message}

	return c.JSON(httpStatus, error)
}
