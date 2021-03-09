package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type User struct {
	ID       uint64    `json:"id"`
	Username string    `json:"username"`
	Products []product `json:"products"`
}

type product struct {
	ID    uint64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/user-products", UserProduct)
	fmt.Printf("user service is up on port: %s", port())
	http.ListenAndServe(port(), nil)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `user service is good`)
}

func UserProduct(w http.ResponseWriter, r *http.Request) {
	p := []product{}
	client := &http.Client{}
	resp, err := client.Get("http://product-service:8100" + "/products")
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	u := User{
		ID:       1,
		Username: "demo@gmail.com",
	}
	u.Products = p
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&u)
}

func port() string {
	p := os.Getenv("USER_SERVICE_PORT")
	if len(strings.TrimSpace(p)) == 0 {
		return ":8081"
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
