package api

import (
	"api/api/db"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	// get all the users in the db
	users, err := db.GetAllUsers()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	// send all the users as response
	return c.JSON(http.StatusOK, users)
}
