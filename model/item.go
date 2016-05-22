package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Item model
type Item struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	CreatedAt time.Time `bson:"createdAt"`
}
