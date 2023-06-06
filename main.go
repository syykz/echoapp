package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"echoapp/interface/controller"
)

func main() {
	e := echo.New()
	e.GET("/users/:id", getUser)


	sample := controller.NewSample(e)
	sample.Handle()

	e.Logger.Fatal(e.Start(":1323"))
}
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

    return c.String(http.StatusOK, id)
}
