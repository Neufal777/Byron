package mysqldb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Byron/core"
	"github.com/Byron/sources"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ttacon/chalk"
)

func SaveArticlesDB() {
	var count int
	db, err := ConnectDb(DB, connectionInformation)

	if err != nil {
		log.Println("no connection")
	}

	defer db.Close()

	articles := sources.ReadArticles("UltimateInventory/General_Collection.json")
	queryInsertArticle, err := db.Prepare("INSERT INTO byronarticles (uniqueid, sourcename, url, downloadurl, title, search,isbn,year,publisher,author,extension,page,language,size,time) VALUES (?, ?, ?, ?, ?, ?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		log.Println("unable to prepare query")
	}
	log.Println("Total articles to insert:", len(articles))
	for i := 0; i < len(articles); i++ {

		alreadyExists := GetArticlesDB("SELECT * FROM byronarticles WHERE url='" + articles[i].Url + "'")

		log.Println("SAME BOOKS:", len(alreadyExists))
		if len(alreadyExists) == 0 {

			fmt.Println(chalk.Green.Color("Inserting article: " + articles[i].UniqueID))

			_, err := queryInsertArticle.Exec(
				articles[i].UniqueID,
				articles[i].SourceName,
				articles[i].Url,
				articles[i].DownloadUrl,
				articles[i].Title,
				articles[i].Search,
				articles[i].Isbn,
				articles[i].Year,
				articles[i].Publisher,
				articles[i].Author,
				articles[i].Extension,
				articles[i].Page,
				articles[i].Language,
				articles[i].Size,
				articles[i].Time,
			)

			count++
			log.Println("Processed:", count)
			if err != nil {
				log.Println("Some error ocurred")
			}

		} else {
			fmt.Println(chalk.Red.Color("Already exists"))
		}

	}
}

func GetArticlesDB(query string) []core.Article {
	var count int
	var ResultArticles []core.Article

	db, err := ConnectDb(DB, connectionInformation)
	if err != nil {
		log.Println("Unable to connect")
	}

	defer db.Close()

	results, err := db.Query(query)

	if err != nil {
		log.Println("Error when fetching")
	}

	for results.Next() {
		art := core.Article{}
		err = results.Scan(
			&art.Id,
			&art.UniqueID,
			&art.SourceName,
			&art.Url,
			&art.DownloadUrl,
			&art.Title,
			&art.Search,
			&art.Isbn,
			&art.Year,
			&art.Publisher,
			&art.Author,
			&art.Extension,
			&art.Page,
			&art.Language,
			&art.Size,
			&art.Time,
		)

		if err != nil {
			log.Println("Mapping issues ")
		}
		ResultArticles = append(ResultArticles, art)
		count++
	}
	fmt.Println(chalk.Green.Color("Total Articles found in DB: " + strconv.Itoa(count)))
	return ResultArticles

}
