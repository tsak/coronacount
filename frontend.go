package main

import (
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

const htmlTemplate = "template.html"

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles(htmlTemplate)
	if err != nil {
		log.WithField("Template", htmlTemplate).WithError(err).Fatal("Unable to parse template")
	}
}

func CoronaCountServer(w http.ResponseWriter, r *http.Request) {
	data := struct{
		Sites []SiteResult
	}{
		Sites: []SiteResult{},
	}

	for _, url := range sites {
		log.WithField("URL", url).Debug("Collecting sites from SiteMap")
		if site, ok := siteMap.Get(url); ok && site.Count > -1 {
			log.WithField("SiteMap", site).Debug("Appending sites to output")
			data.Sites = append(data.Sites, site)
		}
	}

	err := tpl.Execute(w, data)
	if err != nil {
		log.WithError(err).Error("Template error")
	}
}
