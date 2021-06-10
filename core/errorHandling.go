package core

import "strings"

func ErrorsHandling(html string) bool {
	errors := []string{
		"503 Service Temporarily Unavailable",
		"Could not connect to the database",
		"a database error has occurred",
		"504 Gateway Time-out",
		"handshake timeout",
		"403 Forbidden",
	}

	for _, e := range errors {
		if strings.Contains(html, e) {
			return true
		}
	}
	return false
}
