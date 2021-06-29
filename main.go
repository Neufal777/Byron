package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Byron/backend"
	"github.com/Byron/core"
	"github.com/Byron/data"
	"github.com/Byron/sources"
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

	core.TitleFormat("Edward K. Morris Bryan D. Midgley (auth.) Stephen R. Schroeder (eds.) Ecobe", utils.FixUnitedNames("EdwardK.Morris,BryanD.Midgley(auth.),StephenR.Schroeder(eds.)"))
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
		Search:               "AllLibgen",
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
