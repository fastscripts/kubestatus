package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
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
	{"successful", "test.page.gohtml", false, "error rendering go template"},
	{"noTemplate", "test.page.gohtml", true, "randering non-existing go template"},
}

var data = []struct {
	MSG string
}{
	{"Hello World"},
	{"Hello World"},
}

// Config Mock
/*
var app struct{
	Config struct {
		Devmode bool
		TemplatePath string
	}
	TemplateCache map[string]*template.Template
}
*/
// test for render function in renderer.go
func TestRender(t *testing.T) {

	pwd, ma := os.Getwd()
	if ma != nil {
		fmt.Println(ma)
		os.Exit(1)
	}
	fmt.Println(pwd)

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
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	renderer := &Renderer{}

	var buf bytes.Buffer
	err := renderer.Render(&buf, "test", data, c)

	assert.NoError(t, err)

	/*

		for _, tt := range testData {
			t.Run(tt.name, func(t *testing.T) {
				r := &Renderer{}
				b := new(bytes.Buffer)
				// new echo context
				c := echo.New().NewContext(nil, nil)
				err := r.Render(b, tt.template, data, c)
				if tt.errorExpected {
					if err == nil {
						t.Error(tt.errorMessage)
					}
				} else {
					if err != nil {
						t.Error(tt.errorMessage)
					}
				}
			})
		}
	*/
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
