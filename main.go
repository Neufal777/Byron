package main

import "github.com/Byron/data"

func main() {

	// category := flag.String("cat", "math", "category")
	// flag.Parse()
	// fmt.Println(chalk.Magenta.Color("processing " + *category))

	// excat := []string{"salsas", "botas"}

	/*
		source := sources.Source{
			SourceName:           "OpenLibra",
			UrlREGEX:             "<div class=.image-cover.> <a title=.[^<]*href=.([^\"']*)",
			IdREGEX:              "<div class=.image-cover.> <a title=.[^<]*href=.([^\"']*)",
			DownloadUrlREGEX:     "><a class=.command-button btn btn-ol-twitter. href=.([^\"']*)",
			TitleREGEX:           "<h1 itemprop=.name. class[^>]*>([^<]*)",
			IsbnREGEX:            "<div class=.image-cover.> <a title=.[^<]*href=.([^\"']*)",
			YearREGEX:            "<td>A.o:</td><td itemprop=.copyrightYear.>([^<]*)",
			PublisherREGEX:       "<td itemprop=.publisher[^<]*<span itemprop=.name.>([^<]*)",
			AuthorREGEX:          "<span itemprop=.author.>([^<]*)",
			ExtensionREGEX:       "pdf",
			PageREGEX:            "<td itemprop=.numberOfPages.>([^<]*)",
			LanguageREGEX:        "<td>Idioma:</td><td>([^<]*)",
			SizeREGEX:            "<td>Tamaño:</td><td>([^<]*)",
			TimeREGEX:            "<div class=.image-cover.> <a title=.[^<]*href=.([^\"']*)",
			CompletePageUrl:      "https://openlibra.com/es/collection/pag/",
			IncompleteArticleUrl: "",
			AllUrls:              nil,
			Search:               "OpenLibraAll",
		}

		source.GetArticles()
	*/
	data.FilesOrganizer("Inventory/")
}
