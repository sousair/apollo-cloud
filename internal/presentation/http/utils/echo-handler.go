package httputils

import "github.com/labstack/echo/v4"

type EchoHttpHandler interface {
	Handle(c echo.Context) error
}
