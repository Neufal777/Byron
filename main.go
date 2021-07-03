package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Byron/backend"
	"github.com/Byron/data"
	"github.com/Byron/executions"
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

	//sources.ProxiesCleaner()
	//parsecore.ProxyScraping("https://www.bookrix.com/_ebook-h-n-s-mine/")
	//executions.ManyBooksExecution()
	//RegexNoResults()
	// start := time.Now()

	// r := new(big.Int)
	// fmt.Println(r.Binomial(1000, 10))

	// mongodb.SearchArticles("machine")

	// elapsed := time.Since(start)
	// log.Printf("Binomial took %s", elapsed)

	for i := 0; i < 65; i++ {
		fmt.Println("go BookRixSingleExecution(" + strconv.Itoa(i*150) + ", " + strconv.Itoa((i*150)+150) + ")")
	}

}
func ParseExecute() {

	executions.BookRixExecution()
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
