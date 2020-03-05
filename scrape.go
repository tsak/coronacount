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

	c.OnHTML("html", func(e *colly.HTMLElement) {
		content := e.DOM.Find("body").Find("h1,h2,h3,h4,h5,h6,h7,p,ul,ol").Text()

		count := 0
		// Quick and dirty
		for _, s := range []string{"Corona", "Covid-19", "SARS-CoV-2"} {
			count += strings.Count(content, s)                  // Count matches of Corona, ...
			count += strings.Count(content, strings.ToUpper(s)) // Count matches of CORONA, ...
			count += strings.Count(content, strings.ToLower(s)) // Count matches of corona, ...
		}

		if !siteMap.UpdateCount(url, count) {
			log.WithFields(log.Fields{
				"URL":   url,
				"Count": count,
			}).Warn("Count not updated")
			return
		}

		log.WithFields(log.Fields{
			"URL":   url,
			"Count": count,
		}).Info("Scraped")
	})

	// Load URL
	err := c.Visit(url)
	if err != nil {
		log.WithField("URL", url).WithError(err).Error("Unable to scrape")
		return
	}
}
