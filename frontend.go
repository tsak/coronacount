package main

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"sync"
)

var frontend *Content

// Content encapsulates a RWMutex, a template and page bytes
type Content struct {
	sync.RWMutex
	tpl  *template.Template
	page []byte
}

// NewContent loads a given template file and returns a content container
func NewContent(path string) *Content {
	tpl, err := template.ParseFiles(path)
	if err != nil {
		log.WithField("Template", path).WithError(err).Fatalf("Unable to process template")
	}
	return &Content{
		tpl: tpl,
	}
}

// Get returns the content container's page data
func (c *Content) Get() []byte {
	c.RLock()
	defer c.RUnlock()
	return c.page
}

// Render updates the content container's page data by rendering its template
func (c *Content) Render(sr []SiteResult) {
	c.Lock()
	data := struct {
		Sites []SiteResult
	}{
		Sites: sr,
	}

	var b bytes.Buffer
	err := c.tpl.Execute(&b, data)
	if err != nil {
		log.WithError(err).Error("Template error")
	}
	c.page = b.Bytes()
	c.Unlock()
}

// CoronaCountServer returns the frontend's page bytes
func CoronaCountServer(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(frontend.Get())
	if err != nil {
		log.WithError(err).Warn("Writing response")
	}
}
