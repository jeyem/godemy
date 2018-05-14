package main

import (
	"github.com/jeyem/videolearn/app"
	"github.com/jeyem/videolearn/config"
)

func main() {
	a := app.New(config.DefaultConfig)
	a.Register("web")
	a.Run()
}
