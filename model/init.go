package model

import (
	"github.com/jeyem/mogo"

	"navaak/navaak/model/video"
)

func Register(db *mogo.DB, s ...string) {
	for _, m := range s {
		switch m {
		case "video":
			db.LoadIndexes(&video.Video{})
			video.Register(db)
		case "donate":
		case "user":
		}

	}
}
