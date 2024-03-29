package executions

import (
	"log"

	"github.com/Byron/parsecore"
)

func ZBooksExecution() {
	//https://zlibrary.to/top-romanticism-books
	categories := []string{
		"top-romanticism-books",
	}

	for i := 0; i < len(categories)-1; i++ {
		log.Println("Running: ", categories[i], "ID num:", i)
		go ZBooksSingleExecution(categories[i])

	}

	ZBooksSingleExecution(categories[398])
}

func ZBooksSingleExecution(urlCat string) {
	source := parsecore.Source{
		SourceName:           "ZBooks",
		UrlREGEX:             "<h3 itemprop=.name.>[^>]*<a href=.([^\"']*)",
		IncompleteArticleUrl: "https://zlibrary.to/",
		DownloadUrlREGEX:     "<a class=.btn btn-primary dlButton addDownloadedBook. href=.([^\"']*)",
		DownloadUrlComplete:  "https://zlibrary.to/",
		TitleREGEX:           "<h1 itemprop=.name.[^>]*>([^<]*)",
		YearREGEX:            "<div class=.property_label.>Año:</div>[^>]*<div class=.property_value.>([^<]*)",
		AuthorREGEX:          "<div class=.authors.>[^>]*>([^<]*)",
		ExtensionREGEX:       "<div class=.property_label.>Archivo:</div>[^>]*<div class=.property_value.>([^<]*)",
		SizeREGEX:            "<div class=.property_label.>Archivo:</div>[^>]*<div class=.property_value.>([^<]*)",
		PageREGEX:            "<div class=.property_value.><span title=.Pages paperback.>([^<]*)",
		LanguageREGEX:        "<div class=.property_label.>Idioma:</div>[^>]*>([^<]*)",
		CompletePageUrlStart: "https://zlibrary.to/" + urlCat + "?page=",
		AllUrls:              nil,
		Search:               "Zbooks" + urlCat,
	}

	source.GetArticles(1, 20)
}
