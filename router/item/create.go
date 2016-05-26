package item

import (
	"net/http"
  "time"

	"github.com/gin-gonic/gin"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
  "github.com/asaskevich/govalidator"

	Model "github.com/hekike/go-learn/model"
	"github.com/hekike/go-learn/lib"
)

type payload struct {
  Name string `json:"name" binding:"required"`
}

// Create router
func Create(c *gin.Context) {
  db := c.MustGet("db").(*mgo.Database)
  ItemCollection := db.C("items")

  var body payload

  // parse body
  errParse := c.BindJSON(&body)
  if errParse != nil {
    log.WithFields(log.Fields{
      "error": errParse.Error(),
    }).Error("Item cannot be created")

    c.JSON(http.StatusBadRequest, &lib.ServerError{
      Error: "Bad Request",
      StatusCode: http.StatusBadRequest,
      Message: "Invalid payload",
    })
    return
  }

  // create item
  item := &Model.Item{
    ID: bson.NewObjectId(),
    Name: body.Name,
    CreatedAt: time.Now(),
  }

  // validate body
  _, errValidation := govalidator.ValidateStruct(item)
  if errValidation != nil {
    log.WithFields(log.Fields{
      "error": errValidation.Error(),
    }).Error("Item cannot be created")

    c.JSON(http.StatusBadRequest, &lib.ServerError{
      Error: "Bad Request",
      StatusCode: http.StatusBadRequest,
      Message: errValidation.Error(),
    })
    return
  }

  // save item
  errPersist := ItemCollection.Insert(item)
  if errPersist != nil {
    log.WithFields(log.Fields{
      "error": errPersist.Error(),
    }).Error("Item cannot be created")

    c.JSON(lib.InternalServerError.StatusCode, lib.InternalServerError)
    return
  }

  c.JSON(http.StatusCreated, item)
}
