package main

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	starCache = cache.New(5*time.Minute, 10*time.Minute)
)

func setStar(keyword, username string) {
	data, ok := starCache.Get(keyword)
	stars := make([]Star, 0, 5)
	if ok {
		stars = data.([]Star)
	}
	stars = append(stars, Star{UserName: username})
	starCache.Set(keyword, stars, cache.DefaultExpiration)
}

func getStars(keyword string) []*Star {
	data, ok := starCache.Get(keyword)
	stars := make([]Star, 0, 5)
	resStars := make([]*Star, 0, 5)
	if ok {
		stars = data.([]Star)
		for i := 0; i < len(stars); i++ {
			resStars = append(resStars, &stars[i])
		}
	}

	return resStars
}
