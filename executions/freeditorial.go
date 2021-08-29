package executions

import (
	"log"
	"strconv"

	"github.com/Byron/parsecore"
)

func FreeditorialExecution() {
	for i := 0; i < 4500; i++ {
		log.Println("Running from:", i, "to", i+150)
		go FreeditorialSingleExecution(i, i+150)
		i += 149
	}

	FreeditorialSingleExecution(4500, 4555)
}

func FreeditorialSingleExecution(start int, end int) {

	source := parsecore.Source{
		SourceName:           "Freeditorial",
		UrlREGEX:             "class=.book__title.><a href=.([^\"']*)",
		IncompleteArticleUrl: "https://freeditorial.com/",
		TitleREGEX:           "<h1 class=.expandbook__title.>([^<]*)",
		AuthorREGEX:          "<h2 class=.expandbook__author.>[^>]*>([^<]*)",
		CompletePageUrlStart: "https://freeditorial.com/es/books/search?page=",
		CompletePageUrlEnd:   "",
		AllUrls:              nil,
		Search:               "freeditorial:start:" + strconv.Itoa(start) + "end:" + strconv.Itoa(end),
	}

	source.GetArticles(start, end)
}
