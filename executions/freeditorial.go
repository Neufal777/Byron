package executions

import (
	"strconv"

	"github.com/Byron/parsecore"
)

func FreeditorialExecution() {
	go FreeditorialSingleExecution(0, 150)
	go FreeditorialSingleExecution(150, 300)
	go FreeditorialSingleExecution(300, 450)
	go FreeditorialSingleExecution(450, 600)
	go FreeditorialSingleExecution(600, 750)
	go FreeditorialSingleExecution(750, 900)
	go FreeditorialSingleExecution(900, 1050)
	go FreeditorialSingleExecution(1050, 1200)
	go FreeditorialSingleExecution(1200, 1350)
	go FreeditorialSingleExecution(1350, 1500)
	go FreeditorialSingleExecution(1500, 1650)
	go FreeditorialSingleExecution(1650, 1800)
	go FreeditorialSingleExecution(1800, 1950)
	go FreeditorialSingleExecution(1950, 2100)
	go FreeditorialSingleExecution(2100, 2250)
	go FreeditorialSingleExecution(2250, 2400)
	go FreeditorialSingleExecution(2400, 2550)
	go FreeditorialSingleExecution(2550, 2700)
	go FreeditorialSingleExecution(2700, 2850)
	go FreeditorialSingleExecution(2850, 3000)
	go FreeditorialSingleExecution(3000, 3150)
	go FreeditorialSingleExecution(3150, 3300)
	go FreeditorialSingleExecution(3300, 3450)
	go FreeditorialSingleExecution(3450, 3600)
	go FreeditorialSingleExecution(3600, 3750)
	go FreeditorialSingleExecution(3750, 3900)
	go FreeditorialSingleExecution(3900, 4050)
	go FreeditorialSingleExecution(4050, 4200)
	go FreeditorialSingleExecution(4200, 4350)
	go FreeditorialSingleExecution(4350, 4500)
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
