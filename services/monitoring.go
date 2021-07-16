package services

import (
	"fmt"
	"net/http"
)

const MONITORING = 2

var SITES = GetSitesFromFile()

func NeedsToMonitoringWebsites(command int) bool {
	return command == MONITORING
}

func MonitoringWebsites() {
	fmt.Println("Monitorando sites")
	for _, site := range SITES {
		MonitoringWebsite(site)
	}
}

func MonitoringWebsite(site string) {
	response, error := http.Get(site)
	if error != nil {
		fmt.Println("Ocorreu um erro", error)
		return
	}
	if SiteIsLoadedWithSuccess(response.StatusCode) {
		fmt.Println(site, "carregou com sucesso")
		RegisterLog(site, true)
	} else {
		fmt.Println(site, "carregou com erro")
		RegisterLog(site, false)
	}
}

func SiteIsLoadedWithSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
