package web

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/jeyem/godemy/model/video"
	"github.com/labstack/echo"
)

func getVideo(c echo.Context) error {
	v, err := video.FindOne(bson.M{
		"slug": c.Param("slug"),
	})
	if err != nil {
		c.Render(404, "404.html", echo.Map{})
	}
	return c.Render(200, "video.html", echo.Map{
		"video": v.Template(),
	})
}

func stream(c echo.Context) error {
	quality := c.QueryParam("quality")
	src := filepath.Join(app.config.VideoPath, c.Param("src"))
	ext := filepath.Ext(src)
	path := strings.Replace(src, ext, "", -1)
	src = path + quality + ext
	println(src, "----------------")
	f, err := os.Open(src)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	return c.Stream(200, "vid.mp4", f)
}
