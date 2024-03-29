package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Byron/backend"
	"github.com/Byron/data"
	"github.com/Byron/executions"
	"github.com/gorilla/mux"
	"github.com/olekukonko/tablewriter"
)

func init() {
	/*
		Setting variables & configuration
	*/

	os.Setenv("MONGO_CONNECTION", "mongodb+srv://naoufaldracco:7BnUTsAzTEIYlfMPv@cluster0.rs85hn9.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0mongodb+srv://naoufaldracco:7BnUTsAzTEIYlfMPv@cluster0.rs85hn9.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	os.Setenv("MONGO_DATABASE", "dracco_library")
	os.Setenv("MONGO_COLLECTION", "library")

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
	default:
		log.Println("Not found:", *execMode)
		WebExecute()
	}
}

func TestingExecute() {}
func ParseExecute() {
	executions.ManyBooksExecution()
	executions.ArchiveOrgExecution()
	executions.ArchiveOrgExecution()
	executions.FreeditorialExecution()
	executions.LibGenExecution01()
	executions.LibrospdfgratismundoExecution()
	executions.OpenlibraExecution()
	executions.ZBooksExecution()
}

func WebExecute() {
	r := mux.NewRouter()

	//loading files from assets folder
	fh := http.FileServer(http.Dir("./web/assets/"))
	r.PathPrefix("/web/assets/").Handler(http.StripPrefix("/web/assets/", fh))

	r.HandleFunc("/search", backend.SearchResults).Methods("GET")
	r.HandleFunc("/", backend.HomeSearchBar)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Println("PORT=> ", port)
	http.ListenAndServe(":"+port, r)
}
