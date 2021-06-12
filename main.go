package main

import (
	"github.com/Byron/sources"
)

func main() {

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

	//data.DeleteDuplicates("Inventory/")
}
