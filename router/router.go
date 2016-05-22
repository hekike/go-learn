package router

import (
	"myapp/router/item"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func Setup(db *mgo.Database) (*gin.Engine) {
	router := gin.Default()

	// middleware to pass db connection
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// api router
	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/items/:itemID", item.GetById)
	}

	return router
}
