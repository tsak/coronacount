package main

import "sync"

// SiteResult represents a single scrape result
type SiteResult struct {
	Name string
	URL string
	Count int
	Total int
}

// SiteMap encapsulates a map of SiteResults with a RWMutex
type SiteMap struct {
	sync.RWMutex
	sites map[string]SiteResult
}

func NewSiteMap() *SiteMap {
	return &SiteMap{
		sites:   make(map[string]SiteResult),
	}
}

func (m *SiteMap) Get(url string) (result SiteResult, ok bool) {
	m.RLock()
	result, ok = m.sites[url]
	m.RUnlock()
	return result, ok
}

func (m *SiteMap) Delete(url string) {
	m.Lock()
	delete(m.sites, url)
	m.Unlock()
}

func (m *SiteMap) Set(url string, siteResult SiteResult) {
	m.Lock()
	m.sites[url] = siteResult
	m.Unlock()
}
