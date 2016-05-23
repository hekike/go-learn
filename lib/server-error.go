package lib

import "gopkg.in/mgo.v2/bson"
import "net/http"
import "errors"

type ServerError struct {
  Error string    `json:"error"`
  StatusCode int  `json:"statusCode"`
  Message string  `json:"message"`
}

var InternalServerError = ServerError{
  Error: "Server Error",
  StatusCode: http.StatusInternalServerError,
  Message: "Internal server error",
}

func IsObjectIdError(id string) (ServerError, error) {
  if (!bson.IsObjectIdHex(id)) {
    return ServerError{
			Error: "Bad Request",
			StatusCode: http.StatusBadRequest,
			Message: "Invalid ObjectId",
		}, errors.New("Invalid ObjectId")
  }

  return ServerError{}, nil
}

func IsNotFoundError(err error) (ServerError, error) {
  if (err.Error() == "not found") {
    return ServerError{
      Error: "Not Found",
      StatusCode: http.StatusNotFound,
      Message: "Resource is missing",
    }, err
  }

  return ServerError{}, nil
}
