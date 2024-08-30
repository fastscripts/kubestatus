package main

import (
	"log"
	"sync"
	"text/template"

	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/config"
	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/kube"
)

type App struct {
	Name          string
	Webserver     *Webserver
	ConfigPath    string
	Config        *config.AppConfig
	Kube          *kube.Kube
	TemplateCache map[string]*template.Template

	ErrorChan     chan error
	ErrorChanDone chan bool
	Wait          *sync.WaitGroup

	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
