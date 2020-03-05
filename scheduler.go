package main

import (
	"log"
	"time"
)

// Scheduler initially triggers a crawl of all given sites concurrently with a two second delay,
// then reverts to crawling all sites after the configured interval
func Scheduler(sites []string, interval int) {
	if len(sites) == 0 {
		log.Fatal("No sites to crawl")
	}

	// Initially scrape all sites at a 2 second interval to build list
	updateTicker := time.NewTicker(2 * time.Second)
	curr := 0

	for {
		select {
		case <-updateTicker.C:
			go Scrape(sites[curr])
			curr++
			if curr >= len(sites) {
				// Set duration to defined interval after initial scrape
				updateTicker = time.NewTicker(time.Duration(interval) * time.Second)
				curr = 0
			}
		}
	}
}
