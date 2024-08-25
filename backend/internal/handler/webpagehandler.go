package handler

import (
	"encoding/json"
	"net/http"

	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/kube"
	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/models"
	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context, kube *kube.Kube) error {

	//Testdaten
	status := models.ClusterStatus{
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

	//@TODO: App Variable mit sinnvollem Struckt befüllen
	dataMap := make(map[string]interface{})
	//dataMap["status"] = status

	//Testdaten
	jsonData, err := json.Marshal(status)
	if err != nil {
		return c.Render(http.StatusOK, "home.page.gohtml", nil)
	}

	/* 	status, err := kube.GetStatus()
	   	if err != nil {
	   		data := models.TemplateData{
	   			Error: "This is an error" + err.Error(),
	   		}
	   		return c.Render(http.StatusOK, "home.page.gohtml", data)
	   	}
	   	jsonData, err := json.Marshal(status)
	   	if err != nil {
	   		data := models.TemplateData{
	   			Error: "This is an error" + err.Error(),
	   		}
	   		return c.Render(http.StatusOK, "home.page.gohtml", data)
	   	} */

	dataMap["json"] = string(jsonData)

	data := models.TemplateData{
		Data:    dataMap,
		Warning: "Warning",
	}

	return c.Render(http.StatusOK, "home.page.gohtml", data)
}
