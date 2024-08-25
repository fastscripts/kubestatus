package handler

import (
	"net/http"

	"ext-github.swm.de/SWM/rancher-sources/kubestatus/internal/kube"
	"github.com/labstack/echo/v4"
)

func JsonStatus(c echo.Context, kube *kube.Kube) error {

	status, err := kube.GetStatus()
	if err != nil {
		return c.JSON(http.StatusOK, "{Error: "+err.Error()+"  }")
	}

	return c.JSON(http.StatusOK, status)
}

func NodeCount(c echo.Context, kube *kube.Kube) error {

	status, err := kube.GetNodeCount()
	if err != nil {
		return c.JSON(http.StatusOK, "{Error: "+err.Error()+"  }")
	}

	return c.JSON(http.StatusOK, status)
}
