package inthemiddle

import "github.com/labstack/echo/v4"

// SetDefaultServerHeader set noc-cache as default
func SetDefaultServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		c.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Response().Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		return next(c)
	}
}
