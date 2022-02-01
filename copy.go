package router

import (
	"github.com/labstack/echo/v4"
	"hunting-voucher/handler"
)

type API struct {
	Echo *echo.Echo
}

func (api *API) SetupRouter() {
	api.Echo.GET("/", handler.Welcome)
}
