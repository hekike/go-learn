package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Item model
type Item struct {
	ID        bson.ObjectId 		`bson:"_id,omitempty" json:"id"`
	Name      string 						`json:"name"`
	CreatedAt time.Time 				`bson:"createdAt" json:"createdAt"`
}
