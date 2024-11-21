package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestActualizarUsuarioByUsernameFailFormato(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPatch, "/users/juan", strings.NewReader(""))
	response := httptest.NewRecorder()
	t.Run("Actualizar usuario por username con mal formato", func(t *testing.T) {
		UpdateUserByUsername(response, request)
		got := response.Code
		want := 400
		checkHttpStatus(t, got, want)
	})
}

func TestActualizarUsuarioByUsernameFailNoExiste(t *testing.T) {
	var body = strings.NewReader("{\n    \"documento\": \"39453024\",\n    \"username\": \"JuanM\",\n    \"nombres\": \"juan ignacio\",\n    \"apellidos\": \"mesaglio\",\n    \"genero\": \"M\",\n    \"fechaNacimiento\": \"10/01/1996\"\n}")
	request, _ := http.NewRequest(http.MethodPatch, "/users/JuanM", body)
	response := httptest.NewRecorder()
	t.Run("Actualizar usuario por username que no existe.", func(t *testing.T) {
		UpdateUserByUsername(response, request)
		got := response.Code
		want := 404
		checkHttpStatus(t, got, want)
	})
}
func TestCrearUsuarioFail(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/users", strings.NewReader(""))
	response := httptest.NewRecorder()
	t.Run("Crear usuario con mal formato.", func(t *testing.T) {
		CreateUser(response, request)
		got := response.Code
		want := 400
		checkHttpStatus(t, got, want)
	})
}
func TestEliminarUsuarioByUsernameFail(t *testing.T) {
	request, _ := http.NewRequest(http.MethodDelete, "/users/JuanM", nil)
	response := httptest.NewRecorder()
	t.Run("Eliminar usuario que no existe.", func(t *testing.T) {
		DeleteUserByUsername(response, request)
		got := response.Code
		want := 404 // tiene que ser un 404
		checkHttpStatus(t, got, want)
	})
}
func TestObtenerUsuarioByUsernameFail(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/users/JuanM", nil)
	response := httptest.NewRecorder()
	t.Run("Obtener usuario que no existe.", func(t *testing.T) {
		GetUserByUsername(response, request)
		got := response.Code
		want := 404
		checkHttpStatus(t, got, want)
	})
}
func TestObtenerUsuariosEmpty(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/users", nil)
	response := httptest.NewRecorder()
	t.Run("Obtener usuarios vacios.", func(t *testing.T) {
		GetUsers(response, request)
		got := response.Code
		want := 200
		users := response.Body.String()
		checkHttpStatus(t, got, want)
		s := "[]"
		if users != s {
			t.Errorf("got %s, want %s", s, users)
		}
	})
}

func TestCrearUsuario(t *testing.T) {
	var body = strings.NewReader("{\n    \"documento\": \"39453024\",\n    \"username\": \"JuanM\",\n    \"nombres\": \"juan ignacio\",\n    \"apellidos\": \"mesaglio\",\n    \"genero\": \"M\",\n    \"fechaNacimiento\": \"10/01/1996\"\n}")
	request, _ := http.NewRequest(http.MethodPost, "/users", body)
	response := httptest.NewRecorder()
	t.Run("Crear usuario correctamente.", func(t *testing.T) {
		CreateUser(response, request)
		got := response.Code
		want := 201
		checkHttpStatus(t, got, want)
	})
}
func TestObtenerUsuarios(t *testing.T) {
	userString := "{\n    \"documento\": \"39453024\",\n    \"username\": \"JuanM\",\n    \"nombres\": \"juan ignacio\",\n    \"apellidos\": \"mesaglio\",\n    \"genero\": \"M\",\n    \"fechaNacimiento\": \"10/01/1996\"\n}"
	var body = strings.NewReader(userString)
	createRequest, _ := http.NewRequest(http.MethodPost, "/users", body)
	createResponse := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/users", nil)
	response := httptest.NewRecorder()
	t.Run("Obtener usuarios.", func(t *testing.T) {
		CreateUser(createResponse, createRequest)
		GetUsers(response, request)
		got := response.Code
		want := 200
		users := response.Body.String()
		checkHttpStatus(t, got, want)
		if strings.Contains(users, userString) {
			t.Errorf("got %s, want %s", userString, users)
		}
	})
}
func TestEliminarUsuarioByUsername(t *testing.T) {
	var body = strings.NewReader("{\n    \"documento\": \"39453024\",\n    \"username\": \"JuanM\",\n    \"nombres\": \"juan ignacio\",\n    \"apellidos\": \"mesaglio\",\n    \"genero\": \"M\",\n    \"fechaNacimiento\": \"10/01/1996\"\n}")
	createRequest, _ := http.NewRequest(http.MethodPost, "/users", body)
	createResponse := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/users/JuanM", nil)
	response := httptest.NewRecorder()
	t.Run("Eliminar usuario que existe.", func(t *testing.T) {
		CreateUser(createResponse, createRequest)
		DeleteUserByUsername(response, request)
		got := response.Code
		want := 200
		checkHttpStatus(t, got, want)
	})
}
func checkHttpStatus(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
