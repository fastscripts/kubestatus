package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/config"
)

const appName = "kubestatus"

var app = &App{}

func main() {
	app = initialize()
	run(app)
}

func initialize() *App {

	// Initialize Logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	wg := sync.WaitGroup{}

	// Initialize Config with default values
	appConfig := config.AppConfig{
		KubeAccessType: "incluster",
		KubeConfigPath: ".kube/config",
		TemplatePath:   "./web/app/templates",
		Devmode:        true,
		Port:           8080,
		MetricsPort:    8081,
	}

	// Initialize App
	app := &App{
		Name:          appName,
		ConfigPath:    "./config/configuration.json",
		Config:        &appConfig,
		Wait:          &wg,
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
	}

	// load config from file
	app.Config = &config.AppConfig{}
	app.Config.LoadJSONConfiguration(app.ConfigPath)
	app.Config.LoadENVConfiguration()

	// listen for SIGINT and SIGTERM signals and call the shutdown function
	go app.listenForShutdown()

	// listen for errors on the error channel and log them
	go app.listenForErrors()

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

// listenForErrors listens for errors on the error channel and logs them
func (app *App) listenForErrors() {
	for {
		select {
		case err := <-app.ErrorChan:
			app.ErrorLog.Println(err)
		case <-app.ErrorChanDone:
			return
		}
	}
}

// listenForShutdown listens for SIGINT and SIGTERM signals and calls the shutdown function
func (app *App) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.shutdown()
	os.Exit(0)
}

// shutdown performs cleanup tasks and closes channels
func (app *App) shutdown() {
	// perform an cleanup tasks
	app.InfoLog.Println("Cleanup")

	// block until waitgroup is empty
	app.Wait.Wait()

	// close channels
	app.ErrorChanDone <- true

	app.InfoLog.Println(("closing channels and shutting down..."))

	close(app.ErrorChan)
	close(app.ErrorChanDone)

	// Kontext mit Timeout für das Herunterfahren des Servers erstellen
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	ctx := context.Background()

	app.InfoLog.Println("stopping metrics server...")
	app.Webserver.echoPrometheus.Shutdown(ctx)
	app.InfoLog.Println("metrics server stopped")

	app.InfoLog.Println("stopping webserver...")
	app.Webserver.echoWebserver.Shutdown(ctx)
	app.InfoLog.Println("webserver stopped")

}
