package main

import (
	"sort"
	"sync"
)

// SiteResult represents a single scrape result
type SiteResult struct {
	Name     string
	URL      string
	Count    int
	Previous int
	Total    int
}

// SiteMap encapsulates a map of SiteResults with a RWMutex
type SiteMap struct {
	sync.RWMutex
	sites map[string]SiteResult
}

type byCount []SiteResult

func (s byCount) Len() int {
	return len(s)
}
func (s byCount) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byCount) Less(i, j int) bool {
	return s[i].Count > s[j].Count
}

func NewSiteMap() *SiteMap {
	return &SiteMap{
		sites: make(map[string]SiteResult),
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

func (m *SiteMap) UpdateCount(url string, count int) bool {
	updated := false
	m.Lock()
	if site, ok := m.sites[url]; ok {
		updated = true
		site.Previous = site.Count
		site.Count = count
		m.sites[url] = site
	}
	m.Unlock()
	return updated
}

func (m *SiteMap) All() []SiteResult {
	m.RLock()
	var result []SiteResult
	for _, site := range m.sites {
		if site.Count != -1 {
			result = append(result, site)
		}
	}
	sort.Sort(byCount(result))
	m.RUnlock()
	return result
}

func (m *SiteMap) Urls() (urls []string) {
	m.RLock()
	for url := range m.sites {
		urls = append(urls, url)
	}
	m.RUnlock()
	return urls
}
