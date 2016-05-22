package test

import (
  "time"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  Model "myapp/model"
)

func CreateItem (db *mgo.Database, name string) (Model.Item) {
  ItemCollection := db.C("items")

  itemID := bson.NewObjectId()

  err := ItemCollection.Insert(&Model.Item{
    ID: itemID,
    Name: name,
    CreatedAt: time.Now(),
  })

  if err != nil {
		panic(err)
	}

  item := Model.Item{}
  err = ItemCollection.Find(bson.M{
    "_id": itemID,
  }).One(&item)

  if err != nil {
    panic(err)
  }

  return item
}

func RemoveItemById (db *mgo.Database, itemID bson.ObjectId) {
  ItemCollection := db.C("items")

  err := ItemCollection.RemoveId(itemID)

  if err != nil {
    panic(err)
  }
}
