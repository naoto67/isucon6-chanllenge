package main

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	// keyword: html でキャッシュ
	entryCache = cache.New(5*time.Minute, 10*time.Minute)
)

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
