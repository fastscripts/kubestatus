package main

import (
	"strconv"

	"github.com/labstack/echo-contrib/prometheus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Webserver struct {
	Port           int
	PrometheusPort int
	echoPrometheus *echo.Echo
	echoWebserver  *echo.Echo
}

func NewWebserver() *Webserver {

	ws := &Webserver{}
	ws.Port = 8080
	ws.PrometheusPort = 8081

	// Ccreate Prometheus server and Middleware
	ws.echoPrometheus = echo.New()
	ws.echoPrometheus.HideBanner = true

	prom := prometheus.NewPrometheus("echo", nil)

	ws.echoWebserver = echo.New()

	// Scrape metrics from Main Server
	ws.echoWebserver.Use(prom.HandlerFunc)
	// Setup metrics endpoint at another server
	prom.SetMetricsPath(ws.echoPrometheus)

	ws.echoWebserver.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} method: ${method}, uri: ${uri}, status: ${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))

	ws.echoWebserver.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-CSRF-Token,form:_csrf,cookie:_csrf,query:csrf",
	}))

	// static content
	ws.echoWebserver.Static("/css", "web/assets/css")
	ws.echoWebserver.Static("/img", "web/assets/img")
	ws.echoWebserver.Static("/js", "web/assets/js")
	ws.echoWebserver.Static("/data", "web/assets/data")
	ws.echoWebserver.Static("/vendor", "web/assets/vendor")

	ws.echoWebserver.Static("/webfonts", "web/assets/webfonts")

	// load templates
	ws.echoWebserver.Renderer = &Renderer{}
	return ws

}

func (ws *Webserver) Start() {
	go func() {
		ws.echoPrometheus.Logger.Fatal(ws.echoPrometheus.Start(":" + strconv.Itoa(ws.PrometheusPort)))
	}()
	ws.echoWebserver.Logger.Fatal(ws.echoWebserver.Start(":" + strconv.Itoa(ws.Port)))
}
