package main

import (
	"regexp"
)

func getKeywords() ([]string, error) {
	rows, err := db.Query(`
		SELECT keyword FROM entry ORDER BY CHARACTER_LENGTH(keyword) DESC
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
