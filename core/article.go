package core

import (
	"fmt"
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
	art.Authors = utils.FixUnitedNames(art.Authors)

	//Cleaning content
	art.Author = utils.CleanAuthor(art.Author)
	art.Isbn = utils.CleanIsbn(art.Isbn)
	art.Title = utils.CleanTitle(art.Title)

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

func (art *Article) DisplayInformation() {
	/*
		Displays relevant information about the article
	*/

	fmt.Println("ID: " + art.Id)
	fmt.Println("Unique ID: " + art.UniqueID)
	fmt.Println("Title: " + art.Title)
	fmt.Println("Search: " + art.Search)
	fmt.Println("Language: " + art.Language)
	fmt.Println("Download: " + art.DownloadUrl)
	fmt.Println("Isbn: " + art.Isbn)
	fmt.Println("Year: " + art.Year)
	fmt.Println("Publisher: " + art.Publisher)
	fmt.Println("Authors: " + art.Author)
	fmt.Println("Extension: " + art.Extension)
	fmt.Println("Pages: " + art.Page)
	fmt.Println("Size: " + art.Size)
	fmt.Println("Memory: ", art.FileSize)
	fmt.Println("Time: " + art.Time)
}
