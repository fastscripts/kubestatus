package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Renderer struct {
	templates *template.Template
}

func (renderer *Renderer) Render(w io.Writer, templateName string, data interface{}, c echo.Context) error {

	if app.Config.Devmode {

		fileNamePattern := "*.gohtml"

		templ := template.New("")
		err := filepath.Walk(app.Config.TemplatePath,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Println(err.Error())
					return err
				}
				if !info.IsDir() {
					match, err := filepath.Match(fileNamePattern, info.Name())
					if err != nil {
						fmt.Println(err.Error())
						return err
					}
					if match {
						templ, err = templ.ParseGlob(path)
						if err != nil {
							fmt.Println(err.Error())
							return err
						}
					}
				}
				return nil
			})
		if err != nil {
			fmt.Println(err)
		}

		renderer.templates = template.Must(templ, nil)

		return renderer.templates.ExecuteTemplate(w, templateName, data)
	} else {
		// Template aus Cache laden
		/*
			tc, err := NewTemplateCache("./web/app/templates/")
			if err != nil {
				return err
			}
		*/

		templ, ok := app.TemplateCache[templateName]
		if !ok {
			fmt.Println("The template  ", templateName, " does not exist")
			return fmt.Errorf("the template %s does not exist", templateName)
		}

		return templ.Execute(w, data)
	}

}

func NewTemplateCache(templatePath string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(templatePath + "*.page.gohtml")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println(len(pages), " page templates found")
	// Loop through all the page-level templates
	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Loading template: ", name)
		templ, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		templ, err = templ.ParseGlob(templatePath + "*.layout.gohtml")
		if err != nil {
			return nil, err
		}

		templ, err = templ.ParseGlob(templatePath + "*.partial.gohtml")
		if err != nil {
			return nil, err
		}

		componentsPath := templatePath + "components/"
		fileNamePattern := "*.gohtml"

		err = filepath.Walk(componentsPath,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Println(err.Error())
					return err
				}
				if !info.IsDir() {
					match, err := filepath.Match(fileNamePattern, info.Name())
					if err != nil {
						fmt.Println(err.Error())
						return err
					}
					if match {
						templ, err = templ.ParseGlob(path)
						if err != nil {
							fmt.Println(err.Error())
							return err
						}
					}
				}
				return nil
			})
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		cache[name] = templ
	}
	fmt.Println("Templates loaded")
	return cache, nil
}
