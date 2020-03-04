package main

import (
	"bytes"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

func Scrape(url string) {
	c := colly.NewCollector()

	// Send a nice user agent
	c.UserAgent = "CoronaCount/1.0.0 (+https://github.com/tsak/coronacount)"

	// Store response body for counting
	var body []byte
	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode == 200 {
			body = response.Body
		}
	})

	// Parse site title
	title := url
	c.OnHTML("html", func(e *colly.HTMLElement) {
		title = e.DOM.Find("title").First().Text()
	})

	// Load URL
	err := c.Visit(url)
	if err != nil {
		log.WithField("URL", url).WithError(err).Error("Unable to scrape")
		return
	}

	// Quick and dirty
	count := 0
	count += bytes.Count(body, []byte("Corona"))
	count += bytes.Count(body, []byte("corona"))
	count += bytes.Count(body, []byte("CORONA"))

	siteMap.Set(url, SiteResult{
		Name:  title,
		URL:   url,
		Count: count,
		Total: 0,
	})

	log.WithFields(log.Fields{
		"URL":   url,
		"Count": count,
	}).Info("Scraped")
}
