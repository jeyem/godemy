package site

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Subscribe struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Email       string        `bson:"email"`
	Path        string        `bson:"path"`
	Unsubscribe bool          `bson:"unsubscribe"`
	Created     time.Time     `bson:"created"`
}

func (Subscribe) Meta() []mgo.Index {
	return []mgo.Index{
		{Key: []string{"email"}, Unique: true},
	}
}

func (s *Subscribe) Save() error {
	s.Created = time.Now()
	return db.Create(s)
}
