package main

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	keywordCache = cache.New(5*time.Minute, 10*time.Minute)
)

func addKeyword(key string) {
	data, ok := keywordCache.Get("keywords")
	keywords := []string{}
	if ok {
		keywords = data.([]string)
	}
	keywords = append(keywords, key)
	keywordCache.Set("keywords", keywords, cache.DefaultExpiration)
}

func clearKeywords() {
	keywordCache.Delete("keywords")
}

func initKeywords() {
	keywords, err := getKeywords()
	panicIf(err)
	keywordCache.Set("keywords", keywords, cache.DefaultExpiration)
}
func getKeywordsFromCache() []string {
	data, ok := keywordCache.Get("keywords")
	keywords := []string{}
	if ok {
		keywords = data.([]string)
	} else {
		keywords, err := getKeywords()
		panicIf(err)
		keywordCache.Set("keywords", keywords, cache.DefaultExpiration)
	}
	return keywords
}
