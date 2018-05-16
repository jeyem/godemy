package video

import (
	"time"

	"github.com/labstack/echo"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Video struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Title       string        `bson:"title"`
	Slug        string        `bson:"slug"`
	ChapterNum  int           `bson:"chapter_num"`
	Src         string        `bson:"src"`
	Description string        `bson:"description"`
	Duration    time.Duration `bson:"duration"`
	Tags        []string      `bson:"tags"`
	Qualities   []string      `bson:"qualities"`
	Hits        string        `bson:"hits"`
	Ref         bson.ObjectId `bson:"ref"`
	Created     time.Time     `bson:"created"`
	Updated     time.Time     `bson:"updated"`
}

func (Video) Meta() []mgo.Index {
	return []mgo.Index{}
}

func (v Video) Template() echo.Map {
	return echo.Map{
		"id":          v.ID.Hex(),
		"title":       v.Title,
		"slug":        v.Slug,
		"src":         v.Src,
		"description": v.Description,
		"duration":    v.Duration,
		"tags":        v.Tags,
	}

}

func New(id interface{}) (*Video, error) {
	v := new(Video)
	if err := db.Get(v, id); err != nil {
		return nil, err
	}
	return v, nil
}

func Find(q bson.M, limit, page int) (res []Video) {
	db.Where(q).Paginate(limit, page).Find(&res)
	return
}

func FindOne(q bson.M) (*Video, error) {
	v := new(Video)
	if err := db.Where(q).Find(v); err != nil {
		return nil, err
	}
	return v, nil
}
