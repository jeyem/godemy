package model

import (
	"github.com/jeyem/mogo"

	"github.com/jeyem/godemy/model/site"
	"github.com/jeyem/godemy/model/video"
)

func Register(db *mogo.DB, s ...string) {
	for _, m := range s {
		switch m {
		case "video":
			db.LoadIndexes(&video.Video{})
			video.Register(db)
		case "site":
			db.LoadIndexes(&site.Subscribe{})
			site.Register(db)
		case "donate":
		case "user":
		}

	}
}
