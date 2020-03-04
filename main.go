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

	var sitesTxt string
	flag.StringVar(&sitesTxt, "s", "sites.txt", "File to load URLs from")

	var listen string
	flag.StringVar(&listen, "l", "localhost:8080", "Address and port to listen and serve on")

	flag.Parse()

	// Logging
	if debugMode {
		log.SetLevel(log.DebugLevel)
	}
	log.ErrorKey = "Error"

	// Load sites
	sites = LoadSites(sitesTxt)
	log.WithField("URLs", sites).Info("Sites list loaded")

	// Initialise empty, global site map
	siteMap = NewSiteMap()
	for _, url := range sites {
		siteMap.Set(url, SiteResult{
			Name:  url,
			URL:   url,
			Count: -1,
			Total: -1,
		})
	}

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
