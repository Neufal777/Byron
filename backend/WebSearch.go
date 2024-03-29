package backend

import (
	"net/http"
	"strings"

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

	print(search)
	search = strings.TrimSpace(search)
	//art := mongodb.SearchArticles(search)

	searchResults := Search{
		Search: search,
		//Articles: art,
		//Results:  len(art),
	}
	utils.Render(w, "web/templates/search.html", searchResults)
}
