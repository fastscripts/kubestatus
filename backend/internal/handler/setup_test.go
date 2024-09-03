package handler

import (
	"html/template"
	"io"
	"testing"

	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

var e *echo.Echo

// Mock für kube.Kube
type MockKube struct {
	mock.Mock
}

func (m *MockKube) GetStatus() (models.ClusterStatus, error) {
	args := m.Called()
	return args.Get(0).(models.ClusterStatus), args.Error(1)
}

// TemplateRenderer implementiert echo.Renderer für Tests
type TemplateRenderer struct {
	templates *template.Template
}

// Render rendert ein Template
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func TestMain(m *testing.M) {

	e := echo.New()

	e.Renderer = &TemplateRenderer{}

}
