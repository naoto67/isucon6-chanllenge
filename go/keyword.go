package main

import (
	"regexp"
)

func getKeywords() ([]string, error) {
	rows, err := db.Query(`
		SELECT keyword FROM entry
	`)
	panicIf(err)
	keywords := make([]string, 0, 500)
	for rows.Next() {
		var keyword string
		err := rows.Scan(&keyword)
		if err != nil {
			return keywords, err
		}
		keywords = append(keywords, regexp.QuoteMeta(keyword))
	}
	rows.Close()

	return keywords, nil
}

func getKeywordsAndCount(keywords []string) (int, error) {
	rows, err := db.Query(`
		SELECT keyword FROM entry
	`)
	panicIf(err)
	var count int
	for rows.Next() {
		var keyword string
		err := rows.Scan(&keyword)
		if err != nil {
			return count, err
		}
		keywords = append(keywords, regexp.QuoteMeta(keyword))
		count++
	}
	return count, nil
}
