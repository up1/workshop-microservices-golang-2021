package api

import "github.com/labstack/echo/v4"

func MainGroup(e *echo.Echo) {
	e.GET("/users", GetUsers)
}
