package main

import ()

func getKeywords() ([]string, error) {
	rows, err := db.Query(`
		SELECT keyword FROM entry ORDER BY len DESC
	`)
	panicIf(err)
	keywords := make([]string, 0, 8000)
	for rows.Next() {
		var keyword string
		err := rows.Scan(&keyword)
		if err != nil {
			return keywords, err
		}
		keywords = append(keywords, keyword)
	}
	rows.Close()

	return keywords, nil
}
