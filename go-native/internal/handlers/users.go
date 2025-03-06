package handlers

import (
	"bapi/go-native/internal/models"
	"encoding/json"
	"net/http"
	"strings"
)

var users = make(map[string]models.User)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userList := make([]models.User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userList)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users[user.Username] = user
	w.WriteHeader(http.StatusOK)
}

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimPrefix(r.URL.Path, "/users/")
	user, exists := users[username]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimPrefix(r.URL.Path, "/users/")
	if _, exists := users[username]; !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Username = username
	users[username] = user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimPrefix(r.URL.Path, "/users/")
	if _, exists := users[username]; !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(users, username)
	w.WriteHeader(http.StatusOK)
}
