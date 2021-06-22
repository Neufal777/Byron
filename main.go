package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	webbackend "github.com/Byron/WebBackend"
	"github.com/Byron/data"
	"github.com/Byron/mongodb"
	"github.com/Byron/sources"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

}

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

	// i := core.Article{
	// 	SourceName:  "LIBGEN",
	// 	Id:          "4654654",
	// 	UniqueID:    "8796546548975456",
	// 	Url:         "www.github.com",
	// 	Search:      "don't know",
	// 	DownloadUrl: "XXXX",
	// 	Title:       "XXXX",
	// 	Isbn:        "XXXX",
	// 	Year:        "XXXX",
	// 	Publisher:   "XXXX",
	// 	Author:      "XXXX",
	// 	Extension:   "XXXX",
	// 	Page:        "XXXX",
	// 	Language:    "XXXX",
	// 	Size:        "XXXX",
	// 	Time:        "XXXX",
	// }

	// if !mongodb.MongoSearchByUrl("https://libgen.is/book/index.php?md5=5FCA92F369390446BD6DF90380830DA7") {
	// 	log.Println("no existe")
	// }

	log.Println(
		mongodb.SearchArticles(" Doug Lennox - Now You Know Extreme Weather:  The Little Book of Answers"),
	)
}
func ParseExecute() {
	z := sources.Source{
		SourceName:           "LIBGEN",
		UrlREGEX:             "<a href=.book/index.php.md5=([^\"']*)",
		IdREGEX:              "ID:</font></nobr></td><td>([^<]*)",
		DownloadUrlREGEX:     "align=.center.><a href=.([^\"']*). title=.Gen.lib.rus.ec.",
		TitleREGEX:           "<title>Library Genesis:([^<]*)",
		IsbnREGEX:            "ISBN:</font></td><td>([^<]*)",
		YearREGEX:            "Year:</font></nobr></td><td>([^<]*)",
		PublisherREGEX:       "Publisher:</font></nobr></td><td>([^<]*)",
		AuthorREGEX:          "Author.s.:</font></nobr></td><td colspan=3><b>([^<]*)",
		ExtensionREGEX:       "Extension:</font></nobr></td><td>([^<]*)",
		PageREGEX:            "Pages .biblio.tech.:</font></nobr></td><td>([^<]*)",
		LanguageREGEX:        "Language:</font></nobr></td><td>([^<]*)",
		SizeREGEX:            "Size:</font></nobr></td><td>([^<]*)",
		TimeREGEX:            "Time modified:</font></nobr></td><td>([^<]*)",
		CompletePageUrlStart: "https://libgen.is/search.php?&res=100&req=science&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
		IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
		AllUrls:              nil,
		Search:               "science",
	}

	z.GetArticles()

}
func WebExecute() {

	r := mux.NewRouter()

	//loading files from assets folder
	fh := http.FileServer(http.Dir("./web/assets/"))
	r.PathPrefix("/web/assets/").Handler(http.StripPrefix("/web/assets/", fh))

	//web handlers
	r.HandleFunc("/search", webbackend.SearchResults).Methods("GET")
	r.HandleFunc("/", webbackend.HomeSearchBar)

	//identify and assign PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Println("PORT=> ", port)
	http.ListenAndServe(":"+port, r)
}
