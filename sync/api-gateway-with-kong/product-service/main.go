package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type product struct {
	ID    uint64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductConfiguration struct {
	Categories []string `json:"categories"`
}

func main() {
	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/products", Products)
	fmt.Printf("product service is up on port: %s", port())
	http.ListenAndServe(port(), nil)
}

func Products(w http.ResponseWriter, r *http.Request) {
	products := []product{
		{
			ID:    1,
			Name:  "User 1",
			Price: 2000000.00,
		},
		{
			ID:    2,
			Name:  "User 2",
			Price: 500.00,
		},
		{
			ID:    3,
			Name:  "User 3",
			Price: 1500000.00,
		},
		{
			ID:    4,
			Name:  "User 4",
			Price: 50000.00,
		},
		{
			ID:    5,
			Name:  "User 5",
			Price: 20000000.00,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&products)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `product service is good`)
}

func port() string {
	p := os.Getenv("PRODUCT_SERVICE_PORT")
	if len(strings.TrimSpace(p)) == 0 {
		return ":8100"
	}
	return fmt.Sprintf(":%s", p)
}

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}
