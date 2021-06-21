package mysqldb

import (
	"net/http"

	"github.com/Byron/core"
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

	art := GetArticlesDB("select * from byronarticles where title LIKE '%" + search + "%' OR isbn LIKE '%" + search + "%'")
	//art := mongodb.SearchArticles(search)

	searchResults := Search{
		Search:   search,
		Articles: art,
		Results:  len(art),
	}
	utils.Render(w, "web/templates/search.html", searchResults)
}

func HomeSearchBar(w http.ResponseWriter, r *http.Request) {

	utils.Render(w, "web/templates/home.html", nil)
}
