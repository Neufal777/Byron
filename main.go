package main

import "github.com/Byron/sources"

func main() {

	source := sources.Source{
		SourceName:           "OpenLibraaa",
		UrlREGEX:             "<div class=.image-cover.> <a title=.[^<]*href=.([^\"']*)",
		IdREGEX:              "",
		DownloadUrlREGEX:     "<a class=.command-button btn btn-ol-twitter. href=.([^\"']*)",
		TitleREGEX:           "<h1 itemprop=.name. class[^>]*>([^<]*)",
		IsbnREGEX:            "",
		YearREGEX:            "<td itemprop=.copyrightYear.>([^<]*)",
		PublisherREGEX:       "<td itemprop=.publisher[^<]*<span itemprop=.name.>([^<]*)",
		AuthorREGEX:          "<span itemprop=.author.>([^<]*)",
		ExtensionREGEX:       "",
		PageREGEX:            "<td itemprop=.numberOfPages.>([^<]*)",
		LanguageREGEX:        "<td>Idioma:</td><td>([^<]*)",
		SizeREGEX:            "<td>Tama.o:</td><td>([^<]*)",
		TimeREGEX:            "",
		CompletePageUrl:      "https://openlibra.com/es/collection/pag/",
		IncompleteArticleUrl: "",
		AllUrls:              nil,
		Search:               "OpenLibraAllaaa",
	}

	source.GetArticles()

	//data.DeleteDuplicates("Inventory/")
}
