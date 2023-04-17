package main

import (
	"net/http"
	"encoding/json"
	
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
		res, err := json.Marshall(data)
		if err != nil {
			return e.Logger.Fatal(err)
		}
		return c.String(http.StatusOK, res)
	})
	e.Logger.Fatal(e.Start(":80"))
}
