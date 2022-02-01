package main

import (
	"github.com/labstack/echo/v4"
	"hunting-voucher/router"
)

func main() {
	e := echo.New()

	api := router.API{
		Echo:           e,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":1323"))
}
