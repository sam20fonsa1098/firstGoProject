package main

import (
	"firstProject/services"
	"fmt"
	"os"
)

const NAME = "Samuel"
const VERSION float32 = 1.1

func main() {
	initializeMenu()
	runMonitoring()
}

func initializeMenu() {
	fmt.Println("Olá", NAME)
	fmt.Println("Este programa está na versão", VERSION)
}

func runMonitoring() {
	for {
		command := getCommand()
		if services.NeedsToShowLogs(command) {
			services.ShowLogs()
		} else if services.NeedsToMonitoringWebsites(command) {
			services.MonitoringWebsites()
		} else {
			stopGame()
		}
	}
}

func getCommand() int {
	fmt.Println("1- Exibir Logs")
	fmt.Println("2- Iniciar Monitoramento")
	fmt.Println("0- Sair do Programa")
	var command int
	fmt.Scan(&command)
	return command
}

func stopGame() {
	fmt.Println("Monitoramento Finalizado")
	os.Exit(0)
}
