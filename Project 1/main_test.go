package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"encoding/json"
	
	"github.com/stretchr/testify/assert"
)

func TestRootRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello, World!", w.Body.String())
}


func TestPingRoute(t *testing.T) {
	router := setupRouter()
	
	w := httptest.NewRecorder()	

	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"message":"pong"}`, w.Body.String())
}


func TestShowRoute(t *testing.T) {
	router := setupRouter()
	
	w := httptest.NewRecorder()	

	req, _ := http.NewRequest("GET", "/show?team=x-men&member=wolverine", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"team":"x-men","member":"wolverine"}`, w.Body.String())
}


func TestGetUserRoute(t *testing.T) {
	router := setupRouter()
	
	w := httptest.NewRecorder()	

	req, _ := http.NewRequest("GET", "/users/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"id":"1", "name":"John"}`, w.Body.String())
}

func TestPostUserRoute(t *testing.T) {
	router := setupRouter()
	
	w := httptest.NewRecorder()	

	// Подготавливаем JSON-тело запроса
	// jsonBody := `{"name":"Tom","email":"tom@mail.com"}`
	// body := strings.NewReader(jsonBody)

	// Подготавливаем map с данными пользователя
	userData := map[string]interface{}{
		"name":  "Tom",
		"email": "tom@mail.com",
	}

	// Кодируем map в JSON
	jsonBody, err := json.Marshal(userData)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем тело запроса из JSON
	body := strings.NewReader(string(jsonBody))


	req, _ := http.NewRequest("POST", "/users", body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.JSONEq(t, `{"id":"1", "name":"Tom","email":"tom@mail.com"}`, w.Body.String())
}

func TestAdminRoute(t *testing.T) {
	router := setupRouter()
	
	w := httptest.NewRecorder()	

	req, _ := http.NewRequest("GET", "/admin/", nil)
	req.Header.Add("Authorization", "Basic am9lOnNlY3JldA==")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello, Admin!", w.Body.String())
}