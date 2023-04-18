package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Account struct {
	Name string `json:"name"`
}

type HomeResponse struct {
	Operational bool
}

type AccountsResponse struct {
	Accounts []*Account `json:"accounts"`
}

func home(c echo.Context) error {
	data := &HomeResponse{
		Operational: true,
	}
	return c.JSON(http.StatusOK, data)
}

func accounts(c echo.Context) error {
	data := &AccountsResponse{
		Accounts: []*Account{&Account{Name: "deggen"}},
	}
	return c.JSON(http.StatusOK, data)
}

func main() {
	e := echo.New()
	e.GET("/", home)
	e.GET("/accounts", accounts)
	e.Logger.Fatal(e.Start(":3000"))
}
