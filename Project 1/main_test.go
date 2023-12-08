package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestUserRoute(t *testing.T) {
	router := setupRouter()
	
	w := httptest.NewRecorder()	

	req, _ := http.NewRequest("POST", "/users/Tom", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"id":1, "name":"Tom"}`, w.Body.String())
}