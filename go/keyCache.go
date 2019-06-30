package main

import (
	"fmt"
	"html"
	"net/http"
	"regexp"
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

func getKeywordFromCache(key string) bool {
	data, ok := keywordCache.Get("keywords")
	keywords := []string{}
	if ok {
		keywords = data.([]string)
		for _, v := range keywords {
			if v == key {
				return true
			}
		}
	}
	return false
}

func getReplacer(r *http.Request) []string {
	keywords := getKeywordsFromCache()
	rep_data := []string{
		`&`, "&amp;",
		`'`, "&#39;",
		`<`, "&lt;",
		`>`, "&gt;",
		`"`, "&#34;",
	}
	for _, v := range keywords {
		v = regexp.QuoteMeta(v)
		u, err := r.URL.Parse(baseUrl.String() + "/keyword/" + pathURIEscape(v))
		panicIf(err)
		link := fmt.Sprintf("<a href=\"%s\">%s</a>", u, html.EscapeString(v))
		rep_data = append(rep_data, v)
		rep_data = append(rep_data, link)
	}
	return rep_data
}
