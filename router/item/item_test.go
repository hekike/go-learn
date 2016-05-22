package item

import (
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/gavv/httpexpect"

  "github.com/hekike/go-learn/lib"
  "github.com/hekike/go-learn/router"
  "github.com/hekike/go-learn/test"
)

func TestItem(t *testing.T) {
  // test db
  session, db := lib.DBConnect(test.TestDbUri)
  defer session.Close()

  // router
  testRouter := router.Setup(db)
  server := httptest.NewServer(testRouter)
  e := httpexpect.New(t, server.URL)

  item := test.CreateItem(db, "My Item")

  // test invalid body
  e.GET("/api/items/1").
    Expect().
    Status(http.StatusBadRequest).JSON().Object().Equal(map[string]interface{}{
      "Error": "Bad Request",
      "Message": "Invalid ObjectId",
      "StatusCode": 400,
    })

  // test not found
  e.GET("/api/items/56912cbfa233d38145b5ef58").
    Expect().
    Status(http.StatusNotFound).JSON().Object().Equal(map[string]interface{}{
      "Error": "Not Found",
      "Message": "Resource is missing",
      "StatusCode": 404,
    })

  // test success get
  e.GET("/api/items/" + item.ID.Hex()).
    Expect().
    Status(http.StatusOK).JSON().Object().Equal(map[string]interface{}{
      "ID": item.ID.Hex(),
      "Name": item.Name,
      "CreatedAt": item.CreatedAt,
    })

  test.RemoveItemById(db, item.ID)
}
