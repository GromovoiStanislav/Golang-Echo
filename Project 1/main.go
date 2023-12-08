package main

import (
    "net/http"
    
    "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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


	e.GET("/show", func(c echo.Context) error {
		team := c.QueryParam("team")
		member := c.QueryParam("member")

		return c.JSON(http.StatusOK, map[string]string{
			"team": team,
			"member": member,
		})
    })


	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")

		responseData := map[string]interface{}{
			"id":   id,
			"name": "John",
		}
		
		return c.JSON(http.StatusOK, responseData)
	})


	e.POST("/users", func(c echo.Context) error {
		type User struct {
			Id  string `json:"id" xml:"id"`
			Name  string `json:"name" xml:"name" form:"name" query:"name"`
			Email string `json:"email" xml:"email" form:"email" query:"email"`
		}

		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		u.Id = "1"
	
		return c.JSON(http.StatusCreated, u)
		// or
		//return c.XML(http.StatusCreated, u)
	})


	// Group level middleware
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	if username == "joe" && password == "secret" {
		return true, nil
	}
		return false, nil
	}))

	g.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, Admin!")
    })


	return e
}

func main() {
	e := setupRouter()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":8080"))
}