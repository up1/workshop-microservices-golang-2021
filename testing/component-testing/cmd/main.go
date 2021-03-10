package main

import (
	"api/router"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the server")
	e := router.New()
	e.Start(":8080")
}
