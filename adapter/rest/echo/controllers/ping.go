package controllers

import (
	"github.com/labstack/echo/v4"
)

type PingController struct {
	Router *echo.Echo
}

func NewPingController(router *echo.Echo) *PingController {
	return &PingController{
		router,
	}
}

func (thisPC *PingController) Control() {
	thisPC.Router.GET("/ping", func(c echo.Context) (err error) {
		return c.JSON(200, map[string]interface{}{
			"message": "pong",
		})
	})
}

