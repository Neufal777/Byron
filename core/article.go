package core

import (
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
}

func (art *Article) FormatNewArticle() *Article {
	/*
		Format the article
		1 - multiple: authors ISBN and language
	*/
	art.Authors = art.Author
	art.Authors = utils.FixUnitedNames(art.Authors)
	art.Author = utils.FixUnitedNames(art.Author)

	art.Author = strings.Replace(art.Author, ",", ", ", -1)
	art.Author = strings.Replace(art.Author, ";", ", ", -1)
	art.Author = strings.Replace(art.Author, ".", ". ", -1)

	art.Isbn = strings.Replace(art.Isbn, ",", ", ", -1)
	art.Isbn = strings.Replace(art.Isbn, ";", ", ", -1)

	art.Title = strings.Replace(art.Title, ";", "; ", -1)
	art.Title = strings.Replace(art.Title, ":", ": ", -1)
	art.Title = strings.Replace(art.Title, "&nbsp;", " ", -1)

	if strings.Contains(art.Author, ";") {
		Author := strings.Split(art.Author, ";")
		art.Author = Author[0]
	}

	if art.SourceName == "" {
		art.SourceName = "Libgen"
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
	art.Title = strings.TrimSpace(art.Title)
	art.Author = strings.TrimSpace(art.Author)
	art.Authors = strings.TrimSpace(art.Authors)

	art.UniqueID = utils.GetMD5Hash(art.Id + art.Title + art.Url + art.DownloadUrl)

	return art
}

func TitleFormat(title string, authors string) string {
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
