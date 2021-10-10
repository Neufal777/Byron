package executions

import (
	"log"
	"strconv"

	"github.com/Byron/parsecore"
)

func BookRixExecution() {
	for i := 0; i < 9600; i++ {
		log.Println("Running from:", i, "to", i+150)
		go BookRixSingleExecution(i, i+150)
		i += 149
	}

	BookRixSingleExecution(9600, 9750)
}

func BookRixSingleExecution(start int, end int) {

	source := parsecore.Source{
		SourceName:           "BookRix",
		UrlREGEX:             "<big class=.item-title.><a class=.word-break. href=.([^\"']*)",
		IdREGEX:              "",
		DownloadUrlREGEX:     "<a class=.button blue withIcon read. [^>]*href=.([^\"']*)",
		DownloadUrlComplete:  "https://www.bookrix.com",
		TitleREGEX:           "<h2 class=.break-word.>([^<]*)",
		AuthorREGEX:          "<a rel=.author. href=[^>]*>([^<]*)",
		CompletePageUrlStart: "https://www.bookrix.com/books;page:",
		CompletePageUrlEnd:   ".html",
		IncompleteArticleUrl: "https://www.bookrix.com",
		AllUrls:              nil,
		Search:               "BookRix:start:" + strconv.Itoa(start) + "end:" + strconv.Itoa(end),
	}

	source.GetArticles(start, end)
}
