package executions

import (
	"log"
	"strconv"

	"github.com/Byron/parsecore"
)

func OpenlibraExecution() {
	for i := 0; i < 140; i++ {
		log.Println("Running from:", i, "to", i+10)
		go OpenlibraSingleExecution(i, i+10)
		i += 9
	}

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
