package main

import (
	"github.com/jeyem/godemy/app"
	"github.com/jeyem/godemy/config"
)

func main() {
	a := app.New(config.DefaultConfig)
	a.Register("web")
	a.Run()
}
