package services

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const SITES_FILE_PATH = "./sites.txt"

func GetSitesFromFile() []string {
	var sites []string
	file, error := os.Open(SITES_FILE_PATH)
	if error != nil {
		fmt.Println("Ocorreu um erro", error)
		return sites
	}
	reader := bufio.NewReader(file)
	for {
		line, error := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if error == io.EOF {
			break
		}
		sites = append(sites, line)
	}
	file.Close()
	return sites
}
