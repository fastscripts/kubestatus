package main

import (
	"encoding/json"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_addRoutes(t *testing.T) {

	e := echo.New()
	a := &App{
		Webserver: &Webserver{
			echoWebserver: e,
		},
	}

	a.addRoutes()

	//fmt.Println(helper.PrettyPrint(a.Webserver.echoWebserver.Routes()))
	jsonData, err := json.MarshalIndent(a.Webserver.echoWebserver.Routes(), "", "  ")
	if err != nil {
		t.Error(err)
	}

	assert.Contains(t, string(jsonData), "/index.html", "index.html route not found")

}
