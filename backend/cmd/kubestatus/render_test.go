package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	name          string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"successful", "test", false, "error rendering go template"},
	{"noTemplate", "failtes", true, "randering non-existing go template"},
}

var data = []struct {
	MSG string
}{
	{"Hello World"},
	{"Hello World"},
}

func Test_Render(t *testing.T) {

	appConfig := config.AppConfig{
		KubeAccessType: "incluster",
		KubeConfigPath: ".kube/config",
		TemplatePath:   "../../web/app/templates",
		Devmode:        true,
		Port:           8080,
		MetricsPort:    8081,
	}

	app.Config = &appConfig
	app.Config.Devmode = true

	e := echo.New()

	for _, tt := range testData {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		renderer := &Renderer{}

		var buf bytes.Buffer
		err := renderer.Render(&buf, tt.template, data, c)

		if tt.errorExpected {
			assert.Error(t, err, tt.errorMessage)
			continue
		}
		if !tt.errorExpected {
			assert.NoError(t, err)
		}
	}
}

func Test_NewTemplateCache(t *testing.T) {

	testCahce, err := NewTemplateCache("../../web/app/templates")
	// no error expected
	assert.NoError(t, err)
	// check if cache is not nil
	assert.NotNil(t, testCahce)
	// check if cache is not empty
	assert.Greater(t, len(testCahce), 0, "Template cache shoult contain at least one template")

}
