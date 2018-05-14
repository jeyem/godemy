package web

import (
	"github.com/jeyem/videolearn/config"

	"github.com/labstack/echo"
)

var app *web

type web struct {
	http *echo.Echo
	c    *config.Config
}

func Register(e *echo.Echo, c *config.Config) {
	app = new(web)
	app.c = c
	app.http = e
	app.routes()
}

func (w web) routes() {
	w.http.GET("/", index)
}
