package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/Byron/backend"
	"github.com/Byron/data"
	"github.com/Byron/mongodb"
	"github.com/Byron/parsecore"
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
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	mongodb.SearchArticles("machine")

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

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
