package web

import (
	"github.com/jeyem/godemy/model/video"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func index(c echo.Context) error {
	videos := video.Find(bson.M{}, 10, 1)
	lastVids := []echo.Map{}
	for _, vid := range videos {
		lastVids = append(lastVids, vid.Template())
	}
	return c.Render(200, "index.html", echo.Map{
		"last_videos": lastVids,
	})
}
