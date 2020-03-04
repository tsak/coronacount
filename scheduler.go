package main

import (
	"log"
	"time"
)

func Scheduler(sites []string, interval int) {
	if len(sites) == 0 {
		log.Fatal("No sites to crawl")
	}

	updateTicker := time.NewTicker(time.Duration(interval) * time.Second)
	curr := 0

	for {
		select {
		case <-updateTicker.C:
			go Scrape(sites[curr])
			curr++
			if curr >= len(sites) {
				curr = 0
			}
		}
	}
}
