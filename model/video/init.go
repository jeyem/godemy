package video

import "github.com/jeyem/mogo"

var db *mogo.DB

func Register(d *mogo.DB) {
	db = d
}
