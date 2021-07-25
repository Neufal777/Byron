package core

import (
	"regexp"
	"strings"

	"github.com/Byron/utils"
)

type Size struct {
	Ammount string
	Size    string
}
type Article struct {
	SourceName  string
	Id          string
	UniqueID    string
	Url         string
	Search      string
	DownloadUrl string
	Title       string
	Isbn        string
	Year        string
	Publisher   string
	Author      string
	Authors     string
	Extension   string
	Page        string
	Language    string
	Size        string
	FileSize    Size
	Time        string
	Downloaded  int
}

func (art *Article) FormatNewArticle() *Article {
	/*
		Format the article
		1 - multiple: authors ISBN & LANGUAGE, etc..
	*/

	//Authors
	art.Authors = strings.TrimSpace(art.Author)
	art.Authors = FixUnitedNames(art.Authors)

	//Author
	art.Author = FixUnitedNames(art.Author)
	art.Author = strings.ReplaceAll(art.Author, ",", ", ")
	art.Author = strings.ReplaceAll(art.Author, ";", ", ")
	art.Author = strings.ReplaceAll(art.Author, ".", ". ")
	art.Author = strings.ReplaceAll(art.Author, "  ", " ")

	//Isbn
	art.Isbn = strings.ReplaceAll(art.Isbn, ",", ", ")
	art.Isbn = strings.ReplaceAll(art.Isbn, ";", ", ")
	art.Isbn = strings.ReplaceAll(art.Isbn, "  ", " ")

	//Title
	art.Title = strings.ReplaceAll(art.Title, ";", "; ")
	art.Title = strings.ReplaceAll(art.Title, ":", ": ")
	art.Title = strings.ReplaceAll(art.Title, "&nbsp;", " ")
	art.Title = strings.ReplaceAll(art.Title, "&#8211;", "-")
	art.Title = strings.ReplaceAll(art.Title, "  ", " ")

	art.Language = strings.TrimSpace(art.Language)
	art.Extension = strings.TrimSpace(art.Extension)

	if strings.Contains(art.Author, ";") {
		Author := strings.Split(art.Author, ";")
		art.Author = Author[0]
	}

	if strings.Contains(art.Author, ",") {
		Author := strings.Split(art.Author, ",")
		art.Author = Author[0]
	}

	if strings.Contains(art.Year, ";") {
		years := strings.Split(art.Year, ";")
		art.Year = years[0]
	}

	if strings.Contains(art.Year, ",") {
		years := strings.Split(art.Year, ",")
		art.Year = years[0]
	}

	if strings.Contains(art.Size, " ") {
		memory := strings.Split(art.Size, " ")
		art.FileSize = Size{
			Ammount: memory[0],
			Size:    memory[1],
		}
	}

	art.Title = TitleFormat(art.Title, art.Authors)
	art.Author = strings.TrimSpace(art.Author)
	art.UniqueID = utils.GetMD5Hash(art.Id + art.Title + art.Url + art.DownloadUrl)

	if art.DownloadUrl == "" {
		art.DownloadUrl = art.Url
	}

	if art.SourceName == "" {
		art.SourceName = "Libgen"
	}

	if art.Id == "" {
		art.Id = art.UniqueID
	}

	return art
}

func TitleFormat(title string, authors string) string {
	/*
		Delete some author names from the title, minimalism purpouse
	*/
	authors = strings.TrimSpace(authors)
	authors = strings.ReplaceAll(authors, ",", " ")
	authors = strings.ReplaceAll(authors, ";", " ")

	authorsList := strings.Split(authors, " ")

	for _, a := range authorsList {
		title = strings.Replace(title, a, "", -1)
	}

	title = strings.Replace(title, "-", "", -1)
	title = strings.Replace(title, ",", " ", -1)
	title = strings.Replace(title, ";", " ", -1)
	title = strings.Replace(title, "   ", " ", -1)
	title = strings.Replace(title, "  ", " ", -1)

	title = strings.TrimSpace(title)

	return title
}

func FixUnitedNames(name string) string {
	var formatted string
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)

	submatchall := re.FindAllString(name, -1)
	for _, element := range submatchall {
		formatted += element + " "
	}

	/*
		Formatting final name
	*/
	formatted = strings.ReplaceAll(formatted, "-", "")
	formatted = strings.ReplaceAll(formatted, " )", ")")
	formatted = strings.ReplaceAll(formatted, "(", " (")
	formatted = strings.ReplaceAll(formatted, "( ", "(")
	formatted = strings.ReplaceAll(formatted, "  ", " ")
	formatted = strings.TrimSpace(formatted)

	return formatted
}
