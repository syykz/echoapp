package main

import (
	"github.com/labstack/echo/v4"
	"echoapp/interface/controller"
	"echoapp/infrastructure/database"
)

func main() {
	e := echo.New()

	sr := database.NewSampleRepository()
	sample := controller.NewSample(e, sr)
	sample.Handle()

	e.Logger.Fatal(e.Start(":1323"))
}
