package main

import (
	"log"

	"github.com/Byron/db"
	"github.com/Byron/utils"
)

func main() {

	/*
		source := sources.Source{
			SourceName:           "BookRix",
			UrlREGEX:             "<big class=.item-title.><a class=.word-break. href=.([^\"']*)",
			DownloadUrlREGEX:     "data-download=.([^\"']*)",
			DownloadUrlComplete:  "https://www.bookrix.com",
			TitleREGEX:           "<h2 class=.break-word.>([^<]*)",
			AuthorREGEX:          "<a rel=.author. href=[^>]*>([^<]*)",
			TimeREGEX:            "Publication Date:.([^<]*)",
			CompletePageUrlStart: "https://www.bookrix.com/books;page:",
			CompletePageUrlEnd:   ".html",
			IncompleteArticleUrl: "https://www.bookrix.com",
			AllUrls:              nil,
			Search:               "BookRixAll",
		}

		source.GetArticles()
	*/
	//data.DeleteDuplicates("Inventory/")

	log.Println(
		utils.PrettyPrintStruct(
			db.GetArticlesDB("select * from byronarticles WHERE SourceName='openlibra'"),
		),
	)
	//db.SaveArticlesDB()
}
