package http_router

import (
	"net/http"
	
	"github.com/labstack/echo/v4"

	"echo-example/model"
)


func root_handler(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Router{
		Name: "echo",
	})
}