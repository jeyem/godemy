package api

import (
	"github.com/jeyem/godemy/config"

	"github.com/labstack/echo"
)

var app *api

type api struct {
	http   *echo.Echo
	config *config.Config
}

func Register(e *echo.Echo, c *config.Config) {
	app = new(api)
	app.config = c
	app.http = e
	app.routes()
}

func (w api) routes() {
	group := w.http.Group("/api")
	group.POST("/subscribe", subscribe)
}
