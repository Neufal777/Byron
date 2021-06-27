package backend

import (
	"net/http"

	"github.com/Byron/utils"
)

func HomeSearchBar(w http.ResponseWriter, r *http.Request) {

	utils.Render(w, "web/templates/home.html", nil)
}
