package main

import (
	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/config"
)

const appName = "kubestatus"

func main() {
	app := initialize()
	run(app)
}

func initialize() *App {
	appConfig := config.AppConfig{
		KubeAccessType: "incluster",
		KubeConfigPath: ".kube/config",
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

	// @todo accesstype wird überschreiben auf null auch wen kein Konfigfile vorhanden ist
	/* 	app.Config.KubeAccessType = "incluster"
	   	var err error
	   	app.Kube, err = kube.NewKube(app.Config.KubeAccessType, app.Config.KubeConfigPath)
	   	if err != nil {
	   		log.Fatalf("Error could not init kubernetes connection stopping application %v\n", err)
	   	} */

	app.Webserver = NewWebserver()
	app.addRoutes()
	return app
}

func run(app *App) {

	app.Webserver.Start()
}
