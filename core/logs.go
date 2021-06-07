package core

import (
	"fmt"
)

func (art *Article) DisplayInformation() {

	/*
		Displays relevant iformation about the article
	*/
	fmt.Println("ID: " + art.Id)
	fmt.Println("Unique ID: " + art.UniqueID)
	fmt.Println("Title: " + art.Title)
	fmt.Println("Search: " + art.Search)
	fmt.Println("Language:: " + art.Language)
	fmt.Println("Download: " + art.DownloadUrl)
	fmt.Println("Isbn: " + art.Isbn)
	fmt.Println("Isbn List", art.IsbnList)
	fmt.Println("Author List", art.AuthorList)
	fmt.Println("Language List", art.LanguageList)
	fmt.Println("Year: " + art.Year)
	fmt.Println("Publisher: " + art.Publisher)
	fmt.Println("Authors: " + art.Author)
	fmt.Println("Extension: " + art.Extension)
	fmt.Println("Pages: " + art.Page)
	fmt.Println("Size: " + art.Size)
	fmt.Println("Memory: ", art.FileSize)
	fmt.Println("Time: " + art.Time)
}
