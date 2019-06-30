package main

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	entryCache = cache.New(5*time.Minute, 10*time.Minute)
)

func setTopPages(entries []*Entry) {
	e := []Entry{}
	for _, v := range entries {
		e = append(e, *v)
	}
	entryCache.Set("topEntries", e, cache.DefaultExpiration)
}

func addPage(entry Entry) []Entry {
	data, ok := entryCache.Get("topEntries")
	var e []Entry
	var newEntries []Entry
	if ok {
		e = data.([]Entry)
		newEntries = append(newEntries, entry)
		newEntries = append(newEntries, e[:9]...)
	} else {
		newEntries = []Entry{}
		newEntries = append(newEntries, entry)
	}
	entryCache.Set("topEntries", newEntries, cache.DefaultExpiration)

	return newEntries
}

func setHtml(keyword, html string) {
	entryCache.Set(keyword, html, cache.DefaultExpiration)
}

func getHtml(keyword string) (string, bool) {
	data, ok := entryCache.Get(keyword)
	if ok {
		return data.(string), ok
	}
	return "", ok
}
