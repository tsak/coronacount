package main

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strings"
)

func Scrape(url string) {
	c := colly.NewCollector()

	// Send a nice user agent
	c.UserAgent = "CoronaCount/1.0.0 (+https://github.com/tsak/coronacount)"

	// Parse site title
	content := ""
	title := url

	c.OnHTML("html", func(e *colly.HTMLElement) {
		title = e.DOM.Find("title").First().Text()

		content = e.DOM.Find("body").Text()
	})

	// Load URL
	err := c.Visit(url)
	if err != nil {
		log.WithField("URL", url).WithError(err).Error("Unable to scrape")
		return
	}

	// Quick and dirty
	count := 0
	for _, s := range []string{"Corona", "Covid-19", "SARS-CoV-2"} {
		count += strings.Count(content, s)                  // Count matches of Corona, ...
		count += strings.Count(content, strings.ToUpper(s)) // Count matches of CORONA, ...
		count += strings.Count(content, strings.ToLower(s)) // Count matches of corona, ...
	}

	result := SiteResult{
		Name:     title,
		URL:      url,
		Count:    count,
		Previous: count, // Set to current count to avoid comparing current with 0 value
		Total:    0,
	}

	// Get previous result if it exists
	if old, ok := siteMap.Get(url); ok {
		result.Previous = old.Count
	}

	siteMap.Set(url, result)

	log.WithFields(log.Fields{
		"URL":   url,
		"Count": count,
	}).Info("Scraped")
}
