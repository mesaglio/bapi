package main

import (
	"bapi/go-native/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handlers.Ping)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetUsers(w, r)
		case http.MethodPost:
			handlers.CreateUser(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetUserByUsername(w, r)
		case http.MethodPatch:
			handlers.UpdateUserByUsername(w, r)
		case http.MethodDelete:
			handlers.DeleteUserByUsername(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
