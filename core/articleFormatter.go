package core

import (
	"strings"

	"github.com/Byron/utils"
)

func (art *Article) FormatNewArticle() *Article {
	/*
		Format the article
		1 - multiple: authors ISBN and language
	*/

	art.AuthorList = strings.Split(art.Author, ",")
	art.IsbnList = strings.Split(art.Isbn, ",")
	art.LanguageList = strings.Split(art.Language, ",")

	memory := strings.Split(art.Size, " ")
	art.FileSize = Size{
		Ammount: memory[0],
		Size:    memory[1],
	}

	art.UniqueID = utils.GetMD5Hash(art.Id + art.Title + art.Url + art.DownloadUrl)
	return art
}
