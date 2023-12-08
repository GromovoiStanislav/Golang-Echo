package main

import (
    "net/http"
    
    "github.com/labstack/echo/v4"
)

func setupRouter() *echo.Echo {
	e := echo.New()
    
	e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	e.POST("/users/:name", func(c echo.Context) error {
		name := c.Param("name")

		responseData := map[string]interface{}{
			"id":   1,
			"name": name,
		}
		
		return c.JSON(http.StatusOK, responseData)
	})

	return e
}

func main() {
	e := setupRouter()
	e.Logger.Fatal(e.Start(":8080"))
}