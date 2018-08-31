package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", getUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(context echo.Context) error {
	id := context.Param("id")
	return context.String(http.StatusOK, id)
}
