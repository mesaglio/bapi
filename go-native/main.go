package main

import (
	"bapi/go-native/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", methodHandler(map[string]http.HandlerFunc{
		http.MethodGet: handlers.Ping,
	}))

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/users" || r.URL.Path == "/users/" {
			switch r.Method {
			case http.MethodGet:
				handlers.GetUsers(w, r)
			case http.MethodPost:
				handlers.CreateUser(w, r)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
			return
		}

		// Handle /users/{username} endpoints
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
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func methodHandler(handlers map[string]http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler, exists := handlers[r.Method]
		if !exists {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}
