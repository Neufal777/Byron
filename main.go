package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Byron/backend"
	"github.com/Byron/data"
	"github.com/Byron/executions"
	"github.com/Byron/mongodb"
	"github.com/gorilla/mux"
	"github.com/ttacon/chalk"
)

func init() {
	os.Setenv("MONGO_CONNECTION", "mongodb+srv://byron:Black_nebula000@byron.dqvrs.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	os.Setenv("MONGO_DATABASE", "byron")
	os.Setenv("MONGO_COLLECTION", "byronArticles")

	/*
		Setting variables
	*/
	log.Println(chalk.Magenta.Color("Setting up variables"))
	log.Println(chalk.Green.Color("Mongo conenction: " + os.Getenv("MONGO_CONNECTION")))
	log.Println(chalk.Green.Color("Mongo database: " + os.Getenv("MONGO_DATABASE")))
	log.Println(chalk.Green.Color("Mongo collection: " + os.Getenv("MONGO_COLLECTION")))

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
	mongodb.SearchArticles("Strong")

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
