package data

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Byron/mongodb"
	"github.com/Byron/parsecore"
	"github.com/Byron/utils"
	"github.com/ttacon/chalk"
)

func InsertArticles() {
	var count int
	articles := parsecore.ReadArticles("UltimateInventory/General_Collection.json")

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(articles), func(i, j int) { articles[i], articles[j] = articles[j], articles[i] })

	for _, art := range articles {
		mongodb.InsertArticle(&art)
		count++
		fmt.Println(chalk.Magenta.Color("processed:" + utils.AnyTypeToString(count)))
	}
}
