package data

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Byron/mongodb"
	"github.com/Byron/sources"
	"github.com/Byron/utils"
	"github.com/ttacon/chalk"
)

func InsertArticles() {
	var count int
	articles := sources.ReadArticles("UltimateInventory/General_Collection.json")

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(articles), func(i, j int) { articles[i], articles[j] = articles[j], articles[i] })

	for _, art := range articles {
		exists := mongodb.RetrieveArticle(art.Url)

		if len(exists) == 0 {
			mongodb.InsertArticle(&art)
			count++
			fmt.Println(chalk.Magenta.Color("processed:" + utils.AnyTypeToString(count)))
		} else {
			fmt.Println(chalk.Magenta.Color("Already exists"))
		}

	}
}
