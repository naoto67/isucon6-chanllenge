package main

import (
	"regexp"
)

func getKeywords() ([]string, error) {
	rows, err := db.Query(`
		SELECT * FROM entry ORDER BY CHARACTER_LENGTH(keyword) DESC
	`)
	panicIf(err)
	keywords := make([]string, 0, 500)
	for rows.Next() {
		e := Entry{}
		err := rows.Scan(&e.ID, &e.AuthorID, &e.Keyword, &e.Description, &e.UpdatedAt, &e.CreatedAt)
		if err != nil {
			return keywords, err
		}
		keywords = append(keywords, regexp.QuoteMeta(e.Keyword))
	}
	rows.Close()

	return keywords, nil
}
