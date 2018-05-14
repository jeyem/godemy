package video

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Video struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Title       string        `bson:"title"`
	Src         string        `bson:"src"`
	Description string        `bson:"description"`
	Duration    time.Duration `bson:"duration"`
	Size        int           `bson:"size"`
	Tags        []string      `bson:"tags"`
	Hits        string        `bson:"hits"`
	Created     time.Time     `bson:"created"`
	Updated     time.Time     `bson:"updated"`
}

func (Video) Meta() []mgo.Index {
	return []mgo.Index{}
}

func Find(q bson.M) (res []Video) {
	db.Where(q).Find(&res)
	return
}

func FindOne(q bson.M) (*Video, error) {
	v := new(Video)
	if err := db.Where(q).Find(v); err != nil {
		return nil, err
	}
	return v, nil
}
