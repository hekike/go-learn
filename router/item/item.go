package item

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	Model "myapp/model"
	"myapp/lib"
)

// GetById router
func GetById(c *gin.Context) {
	itemID := c.Param("itemID")

	// validate object id
	objectIdErrResp, objectIdErr := lib.IsObjectIdError(itemID)
	if objectIdErr != nil {
		log.WithFields(log.Fields{
			"itemID": itemID,
			"error": "Invalid ObjectId",
		}).Debug("Item cannot get by id")
		c.JSON(objectIdErrResp.StatusCode, objectIdErrResp)
		return
	}

	// get item
	db := c.MustGet("db").(*mgo.Database)
	ItemCollection := db.C("items")

	item := Model.Item{}
	err := ItemCollection.Find(bson.M{
		"_id": bson.ObjectIdHex(itemID),
	}).One(&item)

	// error
	if err != nil {

		// not found error
		notFoundErrResp, notFoundErr := lib.IsNotFoundError(err)
		if notFoundErr != nil {
			log.WithFields(log.Fields{
				"itemID": itemID,
			}).Debug("Item not found by id")
			c.JSON(notFoundErrResp.StatusCode, notFoundErrResp)
		} else {
			// fatal error
			log.WithFields(log.Fields{
				"itemID": itemID,
				"error": err.Error(),
			}).Error("Item cannot get by id")
			c.JSON(lib.InternalServerError.StatusCode, lib.InternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, item)
}
