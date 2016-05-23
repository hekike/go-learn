package lib

import "gopkg.in/mgo.v2"

// DBConnect method
func DBConnect(dbUri string) (*mgo.Session, *mgo.Database) {
	session, err := mgo.Dial(dbUri)
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// get name from uri
	db := session.DB("")

	return session, db
}
