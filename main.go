package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Account struct {
	name string
}

type HomeResponse struct {
	Operational bool
}

type AccountsResponse struct {
	Accounts []*Account
}

func home(c echo.Context) error {
	data := &HomeResponse{
		Operational: true,
	}
	return c.JSON(http.StatusOK, data)
}

func accounts(c echo.Context) error {
	data := make(AccountsResponse, 0)
	data = data.append(&Account{"deggen"})
	return c.JSON(http.StatusOK, data)
}

func main() {
	e := echo.New()
	e.GET("/", home)
	e.GET("/accounts", accounts)
	e.Logger.Fatal(e.Start(":3000"))
}
