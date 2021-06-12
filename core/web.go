package core

import (
	"net/http"

	"github.com/Byron/utils"
)

func WebHome(w http.ResponseWriter, r *http.Request) {

	art := []Article{
		{Id: "79873437"},
		{Id: "79873979"},
	}
	utils.Render(w, "web/templates/home.html", art)
}
