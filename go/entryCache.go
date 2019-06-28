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
		if len(e) == 10 {
			for i := 0; i < len(e)-1; i++ {
				newEntries = append(newEntries, e[i])
			}
		} else {
			for i := 0; i < len(e); i++ {
				newEntries = append(newEntries, e[i])
			}
		}
	} else {
		newEntries = []Entry{}
		newEntries = append(newEntries, entry)
	}

	return newEntries
}
