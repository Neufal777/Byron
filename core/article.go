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
	SourceName   string
	Id           string
	UniqueID     string
	Url          string
	Search       string
	DownloadUrl  string
	Title        string
	Isbn         string
	IsbnList     []string
	Year         string
	Publisher    string
	Author       string
	AuthorList   []string
	Extension    string
	Page         string
	Language     string
	LanguageList []string
	Size         string
	FileSize     Size
	Time         string
}

func (art *Article) FormatNewArticle() *Article {
	/*
		Format the article
		1 - multiple: authors ISBN and language
	*/

	art.Author = strings.Replace(art.Author, " ", "", -1)
	art.Isbn = strings.Replace(art.Isbn, " ", "", -1)
	art.AuthorList = strings.Split(art.Author, ",")
	art.IsbnList = strings.Split(art.Isbn, ",")
	art.LanguageList = strings.Split(art.Language, ",")

	if strings.Contains(art.Author, ";") {
		art.AuthorList = strings.Split(art.Author, ";")
	}
	if strings.Contains(art.Isbn, ";") {
		art.IsbnList = strings.Split(art.Isbn, ";")
	}
	if strings.Contains(art.Language, ";") {
		art.LanguageList = strings.Split(art.Language, ";")
	}

	if strings.Contains(art.Size, " ") {
		memory := strings.Split(art.Size, " ")
		art.FileSize = Size{
			Ammount: memory[0],
			Size:    memory[1],
		}
	}

	art.UniqueID = utils.GetMD5Hash(art.Id + art.Title + art.Url + art.DownloadUrl)
	return art
}
