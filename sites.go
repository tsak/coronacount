package main

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func LoadSites(path string) []string {
	sites, err := readLines(path)
	if err != nil {
		log.WithField("File", path).WithError(err).Fatal("Unable to load sites")
	}
	return sites
}
