package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
	// gorm, usecase, presenter, とかもimportする
)

type Sample interface {
	Handle()
}

type SampleController struct {
	e *echo.Echo
}

func NewSample(e *echo.Echo) *SampleController {
	return &SampleController{e: e}
}

func root(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!B")
}

func (sc *SampleController) Handle() {
	sc.e.GET("/", root)
}