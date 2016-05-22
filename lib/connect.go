package lib

import "gopkg.in/mgo.v2"

// DBConnect method
func DBConnect(dbName string) (*mgo.Session, *mgo.Database) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	db := session.DB(dbName)

	return session, db
}
