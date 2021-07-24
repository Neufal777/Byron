package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Byron/backend"
	"github.com/Byron/core"
	"github.com/Byron/data"
	"github.com/Byron/executions"
	"github.com/Byron/parsecore"
	"github.com/gorilla/mux"
	"github.com/olekukonko/tablewriter"
)

func init() {
	/*
		Setting variables & configuration
	*/
	data := [][]string{
		{"MONGO_CONNECTION", os.Getenv("MONGO_CONNECTION")},
		{"MONGO_DATABASE", os.Getenv("MONGO_DATABASE")},
		{"MONGO_COLLECTION", os.Getenv("MONGO_COLLECTION")},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Value"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}

func main() {
	execMode := flag.String("exec", "web", "Select mode of execution")
	flag.Parse()

	switch *execMode {
	case "web":
		WebExecute()
	case "parse":
		ParseExecute()
	case "insert":
		data.InsertArticlesDatabase()
	case "delete":
		data.DeleteDuplicates("Inventory/")
	case "test":
		TestingExecute()
	case "proxy":
		ProxyTesting()
	default:
		log.Println("Not found:", *execMode)
		WebExecute()
	}
}

func ProxyTesting() {
	parsecore.ProxyScraping("https://es.wikipedia.org/wiki/Los_%C3%81ngeles")
}
func TestingExecute() {
	core.DownloadPDF("https://alex.smola.org/drafts/thebook.pdf", "Machinelearning9893.pdf")
	//executions.ArchiveOrgExecution()
	//data.RecoverSource("Inventory/", "libgen.is")

}

func ParseExecute() {

	executions.FreeditorialExecution()

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
