package web

import (
	"github.com/echo-contrib/pongor"
	"github.com/jeyem/godemy/config"

	"github.com/labstack/echo"
)

var app *web

type web struct {
	http   *echo.Echo
	config *config.Config
}

func Register(e *echo.Echo, c *config.Config) {
	app = new(web)
	app.config = c
	app.http = e
	r := pongor.GetRenderer(
		pongor.PongorOption{
			Reload:    app.config.Debug,
			Directory: app.config.Views,
		},
	)
	app.http.Renderer = r
	app.routes()
}

func (w web) routes() {
	w.http.GET("/", index)
	w.http.GET("/video/:slug", getVideo)
	w.http.GET("/stream/:src", stream)
}
