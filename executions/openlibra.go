package executions

import (
	"strconv"

	"github.com/Byron/parsecore"
)

func OpenlibraExecution() {
	go OpenlibraSingleExecution(0, 20)
	go OpenlibraSingleExecution(20, 40)
	go OpenlibraSingleExecution(40, 60)
	go OpenlibraSingleExecution(60, 80)
	go OpenlibraSingleExecution(80, 100)
	go OpenlibraSingleExecution(100, 120)
	go OpenlibraSingleExecution(120, 140)
	OpenlibraSingleExecution(140, 180)

}

func OpenlibraSingleExecution(start int, end int) {

	source := parsecore.Source{
		SourceName:           "OpenLibra",
		UrlREGEX:             "<div class=.image-cover.> <a title=.[^<]*href=.([^\"']*)",
		DownloadUrlREGEX:     "<a class=.command-button btn btn-ol-twitter. href=.([^\"']*)",
		TitleREGEX:           "<h1 itemprop=.name. class[^>]*>([^<]*)",
		YearREGEX:            "<td itemprop=.copyrightYear.>([^<]*)",
		PublisherREGEX:       "<td itemprop=.publisher[^<]*<span itemprop=.name.>([^<]*)",
		AuthorREGEX:          "<span itemprop=.author.>([^<]*)",
		PageREGEX:            "<td itemprop=.numberOfPages.>([^<]*)",
		LanguageREGEX:        "<td>Idioma:</td><td>([^<]*)",
		SizeREGEX:            "<td>Tama.o:</td><td>([^<]*)",
		CompletePageUrlStart: "https://openlibra.com/es/collection/pag/",
		AllUrls:              nil,
		Search:               "Start:" + strconv.Itoa(start) + "OpenLibraGeneral" + strconv.Itoa(end),
	}

	source.GetArticles(start, end)
}
