package db

import (
	"api/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func createConnection() *sql.DB {
	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(5)

	// check the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func GetAllUsers() ([]models.User, error) {
	// create the postgres db connection
	db := createConnection()
	defer db.Close()

	var users []models.User

	sqlStatement := `SELECT * FROM users`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()

	// iterate
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		users = append(users, user)
	}
	return users, err
}
