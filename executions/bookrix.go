package executions

import (
	"strconv"

	"github.com/Byron/parsecore"
)

func BookRixExecution() {
	go BookRixSingleExecution(0, 150)
	go BookRixSingleExecution(150, 300)
	go BookRixSingleExecution(300, 450)
	go BookRixSingleExecution(450, 600)
	go BookRixSingleExecution(600, 750)
	go BookRixSingleExecution(750, 900)
	go BookRixSingleExecution(900, 1050)
	go BookRixSingleExecution(1050, 1200)
	go BookRixSingleExecution(1200, 1350)
	go BookRixSingleExecution(1350, 1500)
	go BookRixSingleExecution(1500, 1650)
	go BookRixSingleExecution(1650, 1800)
	go BookRixSingleExecution(1800, 1950)
	go BookRixSingleExecution(1950, 2100)
	go BookRixSingleExecution(2100, 2250)
	go BookRixSingleExecution(2250, 2400)
	go BookRixSingleExecution(2400, 2550)
	go BookRixSingleExecution(2550, 2700)
	go BookRixSingleExecution(2700, 2850)
	go BookRixSingleExecution(2850, 3000)
	go BookRixSingleExecution(3000, 3150)
	go BookRixSingleExecution(3150, 3300)
	go BookRixSingleExecution(3300, 3450)
	go BookRixSingleExecution(3450, 3600)
	go BookRixSingleExecution(3600, 3750)
	go BookRixSingleExecution(3750, 3900)
	go BookRixSingleExecution(3900, 4050)
	go BookRixSingleExecution(4050, 4200)
	go BookRixSingleExecution(4200, 4350)
	go BookRixSingleExecution(4350, 4500)
	go BookRixSingleExecution(4500, 4650)
	go BookRixSingleExecution(4650, 4800)
	go BookRixSingleExecution(4800, 4950)
	go BookRixSingleExecution(4950, 5100)
	go BookRixSingleExecution(5100, 5250)
	go BookRixSingleExecution(5250, 5400)
	go BookRixSingleExecution(5400, 5550)
	go BookRixSingleExecution(5550, 5700)
	go BookRixSingleExecution(5700, 5850)
	go BookRixSingleExecution(5850, 6000)
	go BookRixSingleExecution(6000, 6150)
	go BookRixSingleExecution(6150, 6300)
	go BookRixSingleExecution(6300, 6450)
	go BookRixSingleExecution(6450, 6600)
	go BookRixSingleExecution(6600, 6750)
	go BookRixSingleExecution(6750, 6900)
	go BookRixSingleExecution(6900, 7050)
	go BookRixSingleExecution(7050, 7200)
	go BookRixSingleExecution(7200, 7350)
	go BookRixSingleExecution(7350, 7500)
	go BookRixSingleExecution(7500, 7650)
	go BookRixSingleExecution(7650, 7800)
	go BookRixSingleExecution(7800, 7950)
	go BookRixSingleExecution(7950, 8100)
	go BookRixSingleExecution(8100, 8250)
	go BookRixSingleExecution(8250, 8400)
	go BookRixSingleExecution(8400, 8550)
	go BookRixSingleExecution(8550, 8700)
	go BookRixSingleExecution(8700, 8850)
	go BookRixSingleExecution(8850, 9000)
	go BookRixSingleExecution(9000, 9150)
	go BookRixSingleExecution(9150, 9300)
	go BookRixSingleExecution(9300, 9450)
	go BookRixSingleExecution(9450, 9600)
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
		Search:               "start:" + strconv.Itoa(start) + "end:" + strconv.Itoa(end),
	}

	source.GetArticles(start, end)
}
