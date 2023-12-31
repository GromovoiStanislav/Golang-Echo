package main

import (
	"bytes"
	"encoding/json"
	"strconv"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestCreateUserRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	// Assuming you have a JSON payload for creating a user
	// jsonStr := []byte(`{"name":"John Doe","age":30}`)
	
	// Создаем объект пользователя
	user := NewUser{
		Name: "John Doe",
		Age:  30,
	}
	// Сериализуем пользователя в JSON
	jsonStr, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

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

	w := httptest.NewRecorder()

	// Assuming you have a user ID and a JSON payload for updating a user
	userID := 1
	//jsonStr := []byte(`{"name":"Updated User","age":35}`)
	


	// Создаем объект пользователя
	user := NewUser{
		Name: "Updated User",
		Age:  35,
	}
	// Сериализуем пользователя в JSON
	jsonStr, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

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

	// Assert the expected JSON response
	expectedResponse := `{"message":"User deleted","userID":1}`
	assert.JSONEq(t, expectedResponse, w.Body.String())

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

	// Assert the expected JSON response
	expectedResponse := `{"error":"User not found"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())

	// Add more assertions based on the expected response body or headers
}
