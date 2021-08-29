package utils

import (
	"regexp"
	"strings"
)

func CleanIsbn(isbn string) string {
	isbn = strings.ReplaceAll(isbn, ",", ", ")
	isbn = strings.ReplaceAll(isbn, ";", ", ")
	isbn = strings.ReplaceAll(isbn, "  ", " ")

	return isbn
}
func CleanTitle(title string) string {
	title = strings.ReplaceAll(title, ";", "; ")
	title = strings.ReplaceAll(title, ":", ": ")
	title = strings.ReplaceAll(title, "&nbsp;", " ")
	title = strings.ReplaceAll(title, "&#8211;", "-")
	title = strings.ReplaceAll(title, "  ", " ")

	return title
}

func CleanAuthor(author string) string {
	author = FixUnitedNames(author)
	author = strings.ReplaceAll(author, ",", ", ")
	author = strings.ReplaceAll(author, ";", ", ")
	author = strings.ReplaceAll(author, ".", ". ")
	author = strings.ReplaceAll(author, "  ", " ")

	return author
}

func CleanName(name string) string {
	name = strings.ReplaceAll(name, "-", "")
	name = strings.ReplaceAll(name, " )", ")")
	name = strings.ReplaceAll(name, "(", " (")
	name = strings.ReplaceAll(name, "( ", "(")
	name = strings.ReplaceAll(name, "  ", " ")
	name = strings.TrimSpace(name)

	return name
}
func FixUnitedNames(name string) string {
	var formatted string
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)

	submatchall := re.FindAllString(name, -1)
	for _, element := range submatchall {
		formatted += element + " "
	}

	return CleanName(formatted)
}
