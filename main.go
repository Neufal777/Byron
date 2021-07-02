package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Byron/backend"
	"github.com/Byron/core"
	"github.com/Byron/data"
	"github.com/Byron/mongodb"
	"github.com/Byron/parsecore"
	"github.com/Byron/utils"
	"github.com/gorilla/mux"
)

func main() {
	execMode := flag.String("exec", "web", "Select mode of execution")
	flag.Parse()

	switch *execMode {
	case "web":
		WebExecute()
	case "parse":
		ParseExecute()
	case "insert":
		data.InsertArticles()
	case "delete":
		data.DeleteDuplicates("Inventory/")
	case "test":
		TestingExecute()
	default:
		log.Println("Not found:", *execMode)
		WebExecute()
	}

}
func TestingExecute() {

	//sources.ProxiesCleaner()
	//parsecore.ProxyScraping("https://libgen.is/search.php?mode=last&view=simple&phrase=0&timefirst=&timelast=&sort=def&sortmode=ASC&page=10")
	//executions.ManyBooksExecution()
	//RegexNoResults()

	//log.Println(utils.PrettyPrintStruct(mongodb.SearchArticles("9780307588357")))

	arts := []core.Article{
		{SourceName: "openlibra", Title: "Hola", Isbn: "2", Url: "www.google.com"},
		{SourceName: "openlibra", Title: "Adeu", Isbn: "3", Url: "www.facebook.com"},
		{SourceName: "openlibra", Title: "Adeu", Isbn: "3", Url: "www.facebook.com"},
		{SourceName: "openlibra", Title: "Adeu", Isbn: "3", Url: "www.facebook.com"},
		{SourceName: "openlibra", Title: "Adeu", Isbn: "3", Url: "www.facebook.com"},
		{SourceName: "openlibra", Title: "Adeu", Isbn: "3", Url: "www.facebook.com"},
		{SourceName: "openlibra", Title: "Adeu", Isbn: "3", Url: "www.facebook.com"},
		{SourceName: "openlibra", Title: "Adeu", Isbn: "3", Url: "www.facebook.com"},
	}
	log.Println(utils.PrettyPrintStruct(mongodb.ArticleDelDuplicates(arts)))

}
func ParseExecute() {

	z := parsecore.Source{
		SourceName:           "BookRix",
		UrlREGEX:             "<big class=.item-title.><a class=.word-break. href=.([^\"']*)",
		IdREGEX:              "",
		DownloadUrlREGEX:     "data-download=.([^\"']*)",
		DownloadUrlComplete:  "https://www.bookrix.com",
		TitleREGEX:           "<h2 class=.break-word.>([^<]*)",
		AuthorREGEX:          "<a rel=.author. href=[^>]*>([^<]*)",
		CompletePageUrlStart: "https://www.bookrix.com/books;page:",
		CompletePageUrlEnd:   ".html",
		IncompleteArticleUrl: "https://www.bookrix.com",
		AllUrls:              nil,
		Search:               "BookRixAllPart2",
	}

	z.GetArticles()

}
func WebExecute() {

	r := mux.NewRouter()

	//loading files from assets folder
	fh := http.FileServer(http.Dir("./web/assets/"))
	r.PathPrefix("/web/assets/").Handler(http.StripPrefix("/web/assets/", fh))

	//web handlers
	r.HandleFunc("/search", backend.SearchResults).Methods("GET")
	r.HandleFunc("/", backend.HomeSearchBar)

	//identify and assign PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Println("PORT=> ", port)
	http.ListenAndServe(":"+port, r)
}
