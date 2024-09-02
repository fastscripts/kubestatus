package handler

import (
	"testing"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func TestMain(m *testing.M) {

	e := echo.New()

	e.Renderer = &TemplateRenderer{}

}
