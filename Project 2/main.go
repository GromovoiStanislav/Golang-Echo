package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// User represents a simple user model
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type NewUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{}

func listUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
	var newUser User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Just for example, assign a temporary ID
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	return c.JSON(http.StatusCreated, newUser)
}

func getUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	var foundUser *User
	for i := range users {
		if users[i].ID == userID {
			foundUser = &users[i]
			break
		}
	}

	if foundUser == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, foundUser)
}

func updateUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	var updatedUser *User
	for i := range users {
		if users[i].ID == userID {
			if err := c.Bind(&users[i]); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			}
			updatedUser = &users[i]
			break
		}
	}

	if updatedUser == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func deleteUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	userIndex := -1
	for i := range users {
		if users[i].ID == userID {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	users = append(users[:userIndex], users[userIndex+1:]...)

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "User deleted", "userID": userID})
}

func setupRouter() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	users := e.Group("/users")
	{
		users.GET("", listUsers)
		users.POST("", createUser)
		users.GET("/:id", getUser)
		users.PUT("/:id", updateUser)
		users.DELETE("/:id", deleteUser)
	}

	return e
}

func main() {
	e := setupRouter()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":8080"))
}
