package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	//e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up the renderer
	//e.Renderer = &TemplateRenderer{}

	// Mock Kube object
	//mockKube := new(MockKube)

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

	// Call the handler
	if assert.NoError(t, HomePage(c, nil)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Warning")
		// JSON-Daten prüfen
		jsonData, _ := json.Marshal(expectedStatus)
		assert.Contains(t, rec.Body.String(), string(jsonData))
	}
}

func TestUsagePage(t *testing.T) {
	//e := echo.New()
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
}
