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

// Sorting
type byCount []SiteResult

func (s byCount) Len() int { return len(s) }

func (s byCount) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less enables sorting of SiteResult's by count (descending) and by Name if counts are the same
func (s byCount) Less(i, j int) bool {
	if s[i].Count == s[j].Count {
		return s[i].Name < s[j].Name
	}
	return s[i].Count > s[j].Count
}

// NewSiteMap initialises a new site map and returns a pointer
func NewSiteMap() *SiteMap {
	return &SiteMap{
		sites: make(map[string]SiteResult),
	}
}

// Get returns a site map entry for a given URL (thread-safe)
func (m *SiteMap) Get(url string) (result SiteResult, ok bool) {
	m.RLock()
	result, ok = m.sites[url]
	m.RUnlock()
	return result, ok
}

// Delete removes a site map entry for a given URL (thread-safe)
func (m *SiteMap) Delete(url string) {
	m.Lock()
	delete(m.sites, url)
	m.Unlock()
}

// Set adds or replaces a site map entry for a given URL (thread-safe)
func (m *SiteMap) Set(url string, siteResult SiteResult) {
	m.Lock()
	m.sites[url] = siteResult
	m.Unlock()
}

// UpdateCount updates a site maps entries count and stores the previous count (thread-safe)
// Returns true/false if the entry to be updated existed or not
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

// All returns all site map entries sorted by count in descending order (thread-safe)
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

// Urls returns a list of URLs from all entries of the site map
func (m *SiteMap) Urls() (urls []string) {
	m.RLock()
	for url := range m.sites {
		urls = append(urls, url)
	}
	m.RUnlock()
	return urls
}
