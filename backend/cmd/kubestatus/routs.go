package main

import (
	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/handler"
	"github.com/labstack/echo/v4"
)

func (a *App) addRoutes() {
	a.Webserver.echoWebserver.GET("/", func(c echo.Context) error {
		return handler.HomePage(c, a.Kube)
	})
	a.Webserver.echoWebserver.GET("/index.html", func(c echo.Context) error {
		return handler.HomePage(c, a.Kube)
	})
	a.Webserver.echoWebserver.GET("/api/v1/status", func(c echo.Context) error {
		return handler.JsonStatus(c, a.Kube)
	})
	a.Webserver.echoWebserver.GET("/api/v1/nodes", func(c echo.Context) error {
		return handler.NodeCount(c, a.Kube)
	})
}
