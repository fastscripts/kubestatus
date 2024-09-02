package handler

import (
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

func TestHomePage(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Renderer für Tests registrieren
	/* 	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("../../web/app/templates/*.gohtml")), // Pfad zu deinen Templates
	} */
	//e.Renderer = renderer
	e.Renderer = &TemplateRenderer{}

	// Call the handler
	if assert.NoError(t, HomePage(c, nil)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Warning")
		// JSON-Daten prüfen
		//jsonData, _ := json.Marshal(expectedStatus)
		//assert.Contains(t, rec.Body.String(), string(jsonData))
	}
}

/*
func TestUsagePage(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Mock Kube object
	mockKube := new(MockKube)

	// Define expected status data
	expectedStatus := models.ClusterStatus{
		CPU: models.Resources{
			Used:     "1",
			Capacity: "4",
		},
		Memory: models.Resources{
			Used:     "1024",
			Capacity: "4096",
		},
		NodeCount: 2,
	}

	// Set mock return values
	mockKube.On("GetStatus").Return(expectedStatus, nil)

	// Call the handler
	if assert.NoError(t, UsagePage(c, nil)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Warning")
		// JSON-Daten prüfen
		jsonData, _ := json.Marshal(expectedStatus)
		assert.Contains(t, rec.Body.String(), string(jsonData))
	}
} */
