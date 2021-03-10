package router

import (
	"api/api"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	api.MainGroup(e)
	return e
}
