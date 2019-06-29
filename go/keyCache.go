package main

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	keywordCache = cache.New(5*time.Minute, 10*time.Minute)
)

func addLatestKeywords(key string) {
	data, ok := keywordCache.Get("latestKeywords")
	keywords := []string{}
	if ok {
		keywords = data.([]string)
	}
	keywords = append(keywords, key)
	keywordCache.Set("latestKeywords", keywords, cache.DefaultExpiration)
}

func clearLatestKeywords() {
	keywordCache.Delete("latestKeywords")
}
func getLatestKeywords() []string {
	data, ok := keywordCache.Get("latestKeywords")
	keywords := []string{}
	if ok {
		keywords = data.([]string)
	}
	return keywords
}
