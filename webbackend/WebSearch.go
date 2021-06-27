package webbackend

import (
	"net/http"

	"github.com/Byron/core"
	"github.com/Byron/mongodb"
	"github.com/Byron/utils"
)

type Search struct {
	Search   string
	Articles []core.Article
	Results  int
}

func SearchResults(w http.ResponseWriter, r *http.Request) {

	keys := r.URL.Query()
	search := keys["search"][0]

	//art := mysqldb.GetArticlesDB("select * from byronarticles where title LIKE '%" + search + "%' OR isbn LIKE '%" + search + "%'")
	art := mongodb.SearchArticles(search)

	searchResults := Search{
		Search:   search,
		Articles: art,
		Results:  len(art),
	}
	utils.Render(w, "web/templates/search.html", searchResults)
}
