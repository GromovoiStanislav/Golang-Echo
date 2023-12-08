package main

import (
	"bytes"
	"strconv"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestCreateUserRoute(t *testing.T) {
	router := setupRouter()

	// Assuming you have a JSON payload for creating a user
	jsonStr := []byte(`{"name":"John Doe","age":30}`)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	// Assert the expected JSON response
	expectedResponse := `{"id":1,"name":"John Doe","age":30}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
	
	// Add more assertions based on the expected response body or headers
}

func TestListUsersRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// Assert the expected JSON response
	expectedResponse := `[{"id":1,"name":"John Doe","age":30}]`
	assert.JSONEq(t, expectedResponse, w.Body.String())
	
	// Add more assertions based on the expected response body or headers
}

func TestGetUserRoute(t *testing.T) {
	router := setupRouter()

	// Assuming you have a user ID
	userID := 1
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users/"+strconv.Itoa(userID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// Assert the expected JSON response

	expectedResponse := `{"id":1,"name":"John Doe","age":30}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
	
	// Add more assertions based on the expected response body or headers
}

func TestUpdateUserRoute(t *testing.T) {
	router := setupRouter()

	// Assuming you have a user ID and a JSON payload for updating a user
	userID := 1
	jsonStr := []byte(`{"name":"Updated User","age":35}`)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(userID), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	
	// Assert the expected JSON response
	expectedResponse := `{"id":1,"name":"Updated User","age":35}`
	assert.JSONEq(t, expectedResponse, w.Body.String())

	// Add more assertions based on the expected response body or headers
}

func TestDeleteUserRoute(t *testing.T) {
	router := setupRouter()

	// Assuming you have a user ID
	userID := 1
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/users/"+strconv.Itoa(userID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// Add more assertions based on the expected response body or headers
}

func TestDeleteUserRouteNotFound(t *testing.T) {
	router := setupRouter()

	// Assuming you have a user ID
	userID := 1
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/users/"+strconv.Itoa(userID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	// Add more assertions based on the expected response body or headers
}
