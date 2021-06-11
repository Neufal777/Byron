package main

import (
	"log"
	"regexp"

	"github.com/Byron/sources"
)

func main() {

	//category := flag.String("cat", "math", "category")
	//flag.Parse()
	//fmt.Println(chalk.Magenta.Color("processing " + *category))

	source := sources.Source{
		SourceName:           "OpenLibra",
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
		Search:               "OpenLibraAll",
	}

	source.GetArticles()

	//TestingRegex()
	//data.FilesOrganizer("Inventory/")
}

func TestingRegex() {

	texto := "</tr><tr><td>Desde:</td><td>28/03/2019</td></tr><tr><td>Tamaño:</td><td>35.06</td></tr><tr><td>Desde:</td><td>28/03/2019</td></tr><tr><td>Tamaño:</td><td>35.06</td>"

	ArticleTitle, _ := regexp.Compile("<td>Tamaño:</td><td>([^<]*)")
	result := ArticleTitle.FindAllString(texto, -1)[1]

	log.Println(result)
}
