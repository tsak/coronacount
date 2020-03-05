package main

import (
	"bufio"
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"io"
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

func LoadSites(path string) *SiteMap {
	sm := NewSiteMap()

	csvfile, err := os.Open(path)
	if err != nil {
		log.WithField("File", path).WithError(err).Fatal("Unable to load sites")
	}

	r := csv.NewReader(csvfile)
	for {
		// Read each fields from csv
		fields, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.WithError(err).Fatal("CSV read")
		}

		if len(fields) != 2 {
			continue
		}

		// Generate initial sitemap from CSV
		sm.Set(fields[1], SiteResult{
			Name:     fields[0],
			URL:      fields[1],
			Count:    -1,
			Previous: -1,
			Total:    0,
		})
	}

	return sm
}
