package main

import (
  "myapp/router"
  "myapp/lib"
	log "github.com/Sirupsen/logrus"
)

func main() {
  session, db := lib.DBConnect("myapp")
	defer session.Close()

  router := router.Setup(db)

	// configure logger
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
	log.SetLevel(log.DebugLevel)

	router.Run()
}
