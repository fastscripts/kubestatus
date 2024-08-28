package main

import (
	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/config"
)

const appName = "kubestatus"

var app = &App{}

func main() {
	app = initialize()
	run(app)
}

func initialize() *App {
	appConfig := config.AppConfig{
		KubeAccessType: "incluster",
		KubeConfigPath: ".kube/config",
		TemplatePath:   "./web/app/templates",
		Devmode:        true,
		Port:           8080,
		MetricsPort:    8081,
	}

	app := &App{
		Name:       appName,
		ConfigPath: "./config/configuration.json",
		Config:     &appConfig,
	}

	// load config from file
	app.Config = &config.AppConfig{}
	app.Config.LoadJSONConfiguration(app.ConfigPath)
	app.Config.LoadENVConfiguration()

	//@toto defaults werden überschrieben
	app.Config.TemplatePath = "./web/app/templates"

	// Initialize Template Cache wenn nicht im Devmode
	/* 	if !app.Config.Devmode {
	   		var err error
	   		app.TemplateCache, err = NewTemplateCache(app.Config.TemplatePath)
	   		if err != nil {
	   			log.Fatalf("Error could not init template cache %v\n", err)
	   		}
	   	} else {
	   		app.TemplateCache = nil
	   	} */
	// if devmode is false the template cache will be initialized in render.go otherwise it is not used
	app.TemplateCache = nil
	// @todo accesstype wird überschreiben auf null auch wenn kein Konfigfile vorhanden ist
	/* 	app.Config.KubeAccessType = "incluster"
	   	var err error
	   	app.Kube, err = kube.NewKube(app.Config.KubeAccessType, app.Config.KubeConfigPath)
	   	if err != nil {
	   		log.Fatalf("Error could not init kubernetes connection stopping application %v\n", err)
	   	} */

	// Initialize Metrics
	InitMetrics()
	// Initialize Webserver
	app.Webserver = NewWebserver()
	app.addRoutes()
	return app
}

func run(app *App) {

	app.Webserver.Start()
}
