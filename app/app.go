package app

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/jeyem/godemy/config"
	"github.com/jeyem/godemy/controller/web"
	"github.com/jeyem/godemy/model"
	"github.com/jeyem/mogo"
	"github.com/labstack/echo"
)

type App struct {
	db     *mogo.DB
	config *config.Config
	http   *echo.Echo
}

func New(c config.Config) *App {
	a := new(App)
	a.config = &c
	a.db = a.dbConnection()
	a.http = echo.New()
	return a
}

func (a *App) Register(controls ...string) {
	a.loadModels("video")
	for _, c := range controls {
		switch c {
		case "web":
			web.Register(a.http, a.config)
		}
	}
}

func (a *App) Run() {
	a.http.Static("static", filepath.Join(a.config.Views, "static"))
	a.http.Logger.Fatal(a.http.Start(fmt.Sprintf(":%d", a.config.Port)))
}

func (a *App) loadModels(models ...string) {
	model.Register(a.db, models...)
}

func (a *App) Config() config.Config {
	if a.config == nil {
		log.Fatal(ErrorAppInit)
	}
	return *a.config
}

func (a *App) dbConnection() *mogo.DB {
	uri := fmt.Sprintf("%s:%d/%s",
		a.config.MongoHost,
		a.config.MongoPort,
		a.config.MongoDB)
	db, err := mogo.Conn(uri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
