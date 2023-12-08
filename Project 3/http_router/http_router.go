package http_router

import (
	"strconv"
	"fmt"

	"github.com/labstack/echo/v4"
)

func EchoNew(port int) {
	e := echo.New()

    e.GET("/", root_handler)
	
	fmt.Println("serving with ECHO on ", port)
	go e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
}