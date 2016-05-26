package item

import (
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/gavv/httpexpect"
  "gopkg.in/mgo.v2/bson"

  "github.com/hekike/go-learn/lib"
  "github.com/hekike/go-learn/router"
  "github.com/hekike/go-learn/test"
)

func TestCreateItem(t *testing.T) {
  // test db
  session, db := lib.DBConnect(test.TestDbUri)
  defer session.Close()

  // router
  testRouter := router.Setup(db)
  server := httptest.NewServer(testRouter)
  e := httpexpect.New(t, server.URL)

  body := map[string]interface{}{
    "name": "myname",
  }

  // test invalid body
  item := e.POST("/api/items").
    WithJSON(body).
    Expect().
    Status(http.StatusCreated).JSON().Object()

  item.Value("id").String()
  item.Value("name").String().Equal("myname")
  item.Value("name").String()

  // cleanup
  itemId := item.Value("id").Raw().(string)
  test.RemoveItemById(db, bson.ObjectIdHex(itemId))
}
