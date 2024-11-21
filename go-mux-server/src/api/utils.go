package api

import (
	"net/http"
	"net/url"
	"strings"
)

func generateUserJSON() string {
	var _usuarios = ""
	for i := 0; i < len(users); i++ {
		_usuarios = _usuarios + toJson(users[i])
	}
	return "[" + _usuarios + "]"
}

func findUserByUsername(username string) (*User, *int) {
	for i := 0; i < len(users); i++ {
		usuario := users[i]
		if usuario.Username == username {
			return usuario, &i
		}
	}
	return nil, nil
}

func getPathParam(w http.ResponseWriter, r *http.Request) *string {
	_url, err := url.Parse(r.URL.Path)
	username := strings.Split(string(_url.Path), "/")[2]
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}
	return &username
}

func removeByUsername(username string) bool {
	_, indice := findUserByUsername(username)
	if indice == nil {
		return false
	}
	users[len(users)-1], users[*indice] = users[*indice], users[len(users)-1]
	users = users[:len(users)-1]
	return true
}
