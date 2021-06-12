package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Byron/core"
	"github.com/gorilla/mux"
)

func main() {

	/*
		source := sources.Source{
			SourceName:           "BookRix",
			UrlREGEX:             "<big class=.item-title.><a class=.word-break. href=.([^\"']*)",
			DownloadUrlREGEX:     "data-download=.([^\"']*)",
			DownloadUrlComplete:  "https://www.bookrix.com",
			TitleREGEX:           "<h2 class=.break-word.>([^<]*)",
			AuthorREGEX:          "<a rel=.author. href=[^>]*>([^<]*)",
			TimeREGEX:            "Publication Date:.([^<]*)",
			CompletePageUrlStart: "https://www.bookrix.com/books;page:",
			CompletePageUrlEnd:   ".html",
			IncompleteArticleUrl: "https://www.bookrix.com",
			AllUrls:              nil,
			Search:               "BookRixAll",
		}

		source.GetArticles()
	*/

	//_ = db.GetArticlesDB("select * from byronarticles WHERE SourceName='openlibra'")

	r := mux.NewRouter()

	//loading files from assets folder
	fh := http.FileServer(http.Dir("./web/assets/"))
	r.PathPrefix("/web/assets/").Handler(http.StripPrefix("/web/assets/", fh))
	r.HandleFunc("/", core.WebHome)

	//identify and assign PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Println("PORT=> ", port)
	http.ListenAndServe(":"+port, r)
}
