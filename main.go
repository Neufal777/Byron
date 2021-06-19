package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Byron/data"
	"github.com/Byron/db"
	"github.com/gorilla/mux"
)

func main() {

	execMode := flag.String("exec", "web", "Select mode of execution")
	flag.Parse()

	switch *execMode {
	case "web":
		WebExecute()
	case "Parse":
		ParseExecute()
	case "insert":
		db.SaveArticlesDB()
	case "delete":
		data.DeleteDuplicates("Inventory/")
	case "test":
		TestingExecute()
	default:
		log.Println("Not found:", *execMode)
	}

}
func TestingExecute() {}
func ParseExecute()   {}
func WebExecute() {

	r := mux.NewRouter()

	//loading files from assets folder
	fh := http.FileServer(http.Dir("./web/assets/"))
	r.PathPrefix("/web/assets/").Handler(http.StripPrefix("/web/assets/", fh))

	//web handlers
	r.HandleFunc("/search/{search}", db.SearchResults).Methods("GET")
	r.HandleFunc("/", db.HomeSearchBar)

	//identify and assign PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Println("PORT=> ", port)
	http.ListenAndServe(":"+port, r)
}
