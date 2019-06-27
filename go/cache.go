package main

import (
	"fmt"
	"html"
	"net/url"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	repCache *cache.Cache
)

func init() {
	repCache = cache.New(5*time.Minute, 10*time.Minute)

	rows, err := db.Query("SELECT keyword from entry")
	panicIf(err)

	var rep_data []string
	for rows.Next() {
		var key string
		err = rows.Scan(&key)
		panicIf(err)
		u, err := url.Parse(baseUrl.String() + "/keyword/" + pathURIEscape(key))
		panicIf(err)
		link := fmt.Sprintf("<a href=\"%s\">%s</a>", u, html.EscapeString(key))
		rep_data = append(rep_data, key)
		rep_data = append(rep_data, link)
	}

	repCache.Set("reps", rep_data, cache.DefaultExpiration)
}

func addKeyword(key string) {
	data, ok := repCache.Get("reps")
	if ok {
		rep_data := data.([]string)
		u, err := url.Parse(baseUrl.String() + "/keyword/" + pathURIEscape(key))
		panicIf(err)
		link := fmt.Sprintf("<a href=\"%s\">%s</a>", u, html.EscapeString(key))
		rep_data = append(rep_data, key)
		rep_data = append(rep_data, link)
		repCache.Set("reps", rep_data, cache.DefaultExpiration)
	}
}
