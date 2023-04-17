package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type JSONresponse struct {
	Nug string
	Gus int 
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		data := &JSONresponse{
			Nug: "hello",
			Gus: 1234,
		}
		return c.JSON(http.StatusOK, data)
	})
	e.Logger.Fatal(e.Start(":3000"))
}
