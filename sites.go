package main

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

// LoadSites loads sites from CSV and returns a new SiteMap
func LoadSites(path string) *SiteMap {
	sm := NewSiteMap()

	// Open CSV
	file, err := os.Open(path)
	if err != nil {
		log.WithField("File", path).WithError(err).Fatal("Unable to load sites")
	}

	// Read CSV
	r := csv.NewReader(file)
	for {
		// Read fields from CSV
		fields, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.WithError(err).Fatal("CSV read")
		}

		// Ignore lines that don't have two entries exactly
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
