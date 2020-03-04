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
	// TODO: Update list of sites occasionally instead of on request
	data := struct {
		Sites []SiteResult
	}{
		Sites: siteMap.All(),
	}

	err := tpl.Execute(w, data)
	if err != nil {
		log.WithError(err).Error("Template error")
	}
}
