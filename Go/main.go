package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type Contact struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"` // Corrigido JSON tag
	Phone string `json:"phone"`
}

type ContactService struct {
	Contacts map[int]Contact
}

func(c *ContactService) Create(w http.ResponseWriter, t= http.Request){
	var contact Contact
	err :=json.NewDecoder(r.Body).Decode(&contact)
}

func main() {
	// service := &ContactService{Contacts: make(map[int]Contact)}
	mux := http.NewServeMux()
	mux.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
