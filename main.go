package main

import (
  "github.com/hekike/go-learn/router"
  "github.com/hekike/go-learn/lib"
	log "github.com/Sirupsen/logrus"
  "github.com/tj/go-config"
)

type Options struct {
  MongoUri  string  `help:"MongoDB uri" from:"env,flag"`
  LogLevel  string  `help:"log level" from:"env,flag"`
}

func main() {
  // config
  options := Options{
    MongoUri: "mongodb://localhost:27017/myapp",
    LogLevel: "debug",
  }
  config.MustResolve(&options)

	// retreive log level
  logLevel, error := log.ParseLevel(options.LogLevel)
  if error != nil {
    panic(error)
  }

  // setup logger
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
	log.SetLevel(logLevel)

  // setup db
  session, db := lib.DBConnect(options.MongoUri)
  defer session.Close()

  // setup router
  router := router.Setup(db)
	router.Run()
}
