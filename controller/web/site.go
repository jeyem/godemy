package web

import "github.com/labstack/echo"

func index(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"message": "Hello!",
	})
}
