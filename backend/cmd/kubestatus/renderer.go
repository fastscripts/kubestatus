package main

import (
	"fmt"
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Renderer struct {
	templates *template.Template
}

func (renderer *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Template bei jedem Aufruf neu laden
	// @todo: nur im dev mode
	// @todo: Template Verzeichnis auslesen

	templ, err := template.New("").ParseGlob("./web/app/templates/*.gohtml")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	templ, err = templ.ParseGlob("./web/app/templates/components/*.gohtml")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	templ, err = templ.ParseGlob("./web/app/templates/components/messages/*.gohtml")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	templ, err = templ.ParseGlob("./web/app/templates/components/cards/*.gohtml")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	renderer.templates = template.Must(templ, nil)

	return renderer.templates.ExecuteTemplate(w, name, data)
}
