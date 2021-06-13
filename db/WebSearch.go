package db

import (
	"net/http"

	"github.com/Byron/core"
	"github.com/Byron/utils"
	"github.com/gorilla/mux"
)

type Search struct {
	Search   string
	Articles []core.Article
	Results  int
}

func SearchResults(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	search := vars["search"]

	art := GetArticlesDB("select * from byronarticles where title LIKE '%" + search + "%'")

	searchResults := Search{
		Search:   vars["search"],
		Articles: art,
		Results:  len(art),
	}
	utils.Render(w, "web/templates/search.html", searchResults)
}
