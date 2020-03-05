package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	Run()
}

var siteMap *SiteMap
var sites []string

func Run() {
	// Command line flags
	var debugMode bool
	flag.BoolVar(&debugMode, "d", false, "Debug mode")

	var interval int
	flag.IntVar(&interval, "i", 300, "Scrape interval in seconds")

	var sitesCsv string
	flag.StringVar(&sitesCsv, "s", "sites.csv", `CSV file to load URLs from, format is "Site title", URL`)

	var listen string
	flag.StringVar(&listen, "l", "localhost:8080", "Address and port to listen and serve on")

	var tplFile string
	flag.StringVar(&tplFile, "t", "template.html", "HTML template")

	flag.Parse()

	// Logging
	if debugMode {
		log.SetLevel(log.DebugLevel)
	}
	log.ErrorKey = "Error"

	// Load sites from CSV
	siteMap = LoadSites(sitesCsv)
	sites = siteMap.Urls()
	log.WithField("URLs", sites).Info("Sites list loaded")

	// Initialise template
	frontend = NewContent(tplFile)
	log.WithField("Template", tplFile).Info("Template initialised")

	// Start scheduler
	go Scheduler(sites, interval)

	log.WithField("Address", listen).Info("Starting service")

	// Start server
	http.HandleFunc("/", CoronaCountServer)
	srv := &http.Server{
		Addr:         listen,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.WithError(err).Error("HTTP server")
	}
}
