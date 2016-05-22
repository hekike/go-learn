package main

import (
  "github.com/hekike/go-learn/router"
  "github.com/hekike/go-learn/lib"
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
