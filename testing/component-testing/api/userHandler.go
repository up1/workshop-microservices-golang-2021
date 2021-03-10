package api

import (
	"api/api/db"
	"api/models"
	"encoding/json"
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

func GetUsersFromAPI(c echo.Context) error {
	// get all the users in the db
	client := &http.Client{}
	resp, err := client.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	defer resp.Body.Close()

	users := []models.User{}
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	// send all the users as response
	return c.JSON(http.StatusOK, users)
}
