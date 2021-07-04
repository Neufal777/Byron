package mongodb

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Byron/core"
	"github.com/Byron/utils"
	"github.com/ttacon/chalk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MongoSearchByUrl(url string) bool {
	client, ctx, err := ConnectMongoDB()

	defer client.Disconnect(ctx)

	if err != nil {
		panic(err)
	}

	byronDatabase := client.Database(os.Getenv("MONGO_DATABASE"))
	byronArticlesCollection := byronDatabase.Collection(os.Getenv("MONGO_COLLECTION"))

	filterCursor, err := byronArticlesCollection.Find(ctx, bson.M{"Url": url})

	if err != nil {
		panic(err)
	}

	var articlesRetrieved []bson.M
	if err = filterCursor.All(ctx, &articlesRetrieved); err != nil {
		log.Fatal(err)
	}

	if len(articlesRetrieved) == 0 {
		return false
	}
	return true
}

func InsertArticle(article *core.Article) {
	client, ctx, err := ConnectMongoDB()

	defer client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}

	byronDatabase := client.Database(os.Getenv("MONGO_DATABASE"))
	byronArticlesCollection := byronDatabase.Collection(os.Getenv("MONGO_COLLECTION"))

	if !MongoSearchByUrl(article.Url) {

		_, err = byronArticlesCollection.InsertOne(ctx, bson.D{
			{Key: "UniqueID", Value: article.UniqueID},
			{Key: "SourceName", Value: article.SourceName},
			{Key: "Url", Value: article.Url},
			{Key: "DownloadUrl", Value: article.DownloadUrl},
			{Key: "Title", Value: article.Title},
			{Key: "Search", Value: article.Search},
			{Key: "Isbn", Value: article.Isbn},
			{Key: "Year", Value: article.Year},
			{Key: "Publisher", Value: article.Publisher},
			{Key: "Author", Value: article.Author},
			{Key: "Extension", Value: article.Extension},
			{Key: "Page", Value: article.Page},
			{Key: "Language", Value: article.Language},
			{Key: "Size", Value: article.Size},
			{Key: "Time", Value: article.Time},
		})

		if err != nil {
			log.Println(err)
		}

		fmt.Println(chalk.Green.Color("Inserted correctly: " + article.UniqueID))

	} else {
		fmt.Println(chalk.Red.Color("This article exists: " + article.UniqueID))
	}
}

func SearchArticles(search string) []core.Article {
	/*
		Get all possible matches and store them :)
	*/
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	var AllArticles []core.Article

	AllArticles = append(AllArticles, GetArticlesRegex(search)...)
	cleanedCollection := ArticleDelDuplicates(AllArticles)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	return cleanedCollection
}

func ArticleDelDuplicates(allArticles []core.Article) []core.Article {

	var all []core.Article
	dups := 0

	arti := make(map[string]core.Article)

	for _, articles := range allArticles {
		_, ok := arti[articles.Url]

		if !ok {
			arti[articles.Url] = articles
		} else {
			dups++
		}
	}

	for _, v := range arti {
		all = append(all, v)
	}

	log.Println("Total Found:", len(all))
	log.Println("Duplicates:", dups)
	return all
}

func GetArticlesRegex(search string) []core.Article {
	client, ctx, err := ConnectMongoDB()

	var (
		byronDatabase           = client.Database(os.Getenv("MONGO_DATABASE"))
		byronArticlesCollection = byronDatabase.Collection(os.Getenv("MONGO_COLLECTION"))
		AllArticles             []core.Article
	)

	if err != nil {
		log.Fatal(err)
	}

	queries := []bson.M{
		{"Title": bson.M{"$regex": primitive.Regex{Pattern: search, Options: "is"}}},
		{"Isbn": bson.M{"$regex": primitive.Regex{Pattern: search, Options: "is"}}},
		{"Author": bson.M{"$regex": primitive.Regex{Pattern: search, Options: "is"}}},
	}

	for _, query := range queries {

		var articlesRetrieved []bson.M

		filterCursor, err := byronArticlesCollection.Find(ctx, query)

		if err != nil {
			log.Panic(err)
		}

		if err = filterCursor.All(ctx, &articlesRetrieved); err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(articlesRetrieved); i++ {

			formatedIsbn := utils.AnyTypeToString(articlesRetrieved[i]["Isbn"])
			formatedIsbn = strings.TrimSpace(formatedIsbn)
			formatedIsbn = strings.ReplaceAll(formatedIsbn, "-", "")

			if strings.Contains(formatedIsbn, ";") {
				allIsbns := strings.Split(formatedIsbn, ";")
				formatedIsbn = allIsbns[0]
			}

			if strings.Contains(formatedIsbn, ",") {
				allIsbns := strings.Split(formatedIsbn, ",")
				formatedIsbn = allIsbns[0]
			}

			/*
				Format and insert only ammount and type of size
			*/

			var formattedAmmount string
			var formattedUnit string

			size := utils.AnyTypeToString(articlesRetrieved[i]["Size"])

			if strings.Contains(size, " ") {
				memory := strings.Split(size, " ")
				formattedAmmount = memory[0]
				formattedUnit = memory[1]
			}

			AllArticles = append(AllArticles, core.Article{
				UniqueID:    utils.AnyTypeToString(articlesRetrieved[i]["UniqueID"]),
				SourceName:  utils.AnyTypeToString(articlesRetrieved[i]["SourceName"]),
				Url:         utils.AnyTypeToString(articlesRetrieved[i]["Url"]),
				DownloadUrl: utils.AnyTypeToString(articlesRetrieved[i]["DownloadUrl"]),
				Title:       utils.AnyTypeToString(articlesRetrieved[i]["Title"]),
				Search:      utils.AnyTypeToString(articlesRetrieved[i]["Search"]),
				Isbn:        formatedIsbn,
				Year:        utils.AnyTypeToString(articlesRetrieved[i]["Year"]),
				Publisher:   utils.AnyTypeToString(articlesRetrieved[i]["Publisher"]),
				Author:      strings.TrimSpace(utils.AnyTypeToString(articlesRetrieved[i]["Author"])),
				Extension:   utils.AnyTypeToString(articlesRetrieved[i]["Extension"]),
				Page:        utils.AnyTypeToString(articlesRetrieved[i]["Page"]),
				Language:    utils.AnyTypeToString(articlesRetrieved[i]["Language"]),
				Size:        utils.AnyTypeToString(articlesRetrieved[i]["Size"]),
				FileSize: core.Size{
					Ammount: formattedAmmount,
					Size:    formattedUnit,
				},
				Time: utils.AnyTypeToString(articlesRetrieved[i]["Time"]),
			})
		}
	}

	fmt.Println(chalk.Green.Color("Articles found by :" + strconv.Itoa(len(AllArticles))))
	return AllArticles
}

// func GETARTICLESTESTINGMONGO(search string) []core.Article {
// 	client, ctx, err := ConnectMongoDB()

// 	var (
// 		byronDatabase           = client.Database(os.Getenv("MONGO_DATABASE"))
// 		byronArticlesCollection = byronDatabase.Collection(os.Getenv("MONGO_COLLECTION"))
// 		AllArticles             []core.Article
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	queries := []bson.M{
// 		{"Title": bson.M{"$regex": primitive.Regex{Pattern: search, Options: "is"}}},
// 		{"Isbn": bson.M{"$regex": primitive.Regex{Pattern: search, Options: "is"}}},
// 		{"Author": bson.M{"$regex": primitive.Regex{Pattern: search, Options: "is"}}},
// 	}

// 		var articlesRetrieved []bson.M

// 		filterCursor, err := byronArticlesCollection.Find(ctx, query)

// 		if err != nil {
// 			log.Panic(err)
// 		}

// 		if err = filterCursor.All(ctx, &articlesRetrieved); err != nil {
// 			log.Fatal(err)
// 		}

// 		for i := 0; i < len(articlesRetrieved); i++ {

// 			formatedIsbn := utils.AnyTypeToString(articlesRetrieved[i]["Isbn"])
// 			formatedIsbn = strings.TrimSpace(formatedIsbn)
// 			formatedIsbn = strings.ReplaceAll(formatedIsbn, "-", "")

// 			if strings.Contains(formatedIsbn, ";") {
// 				allIsbns := strings.Split(formatedIsbn, ";")
// 				formatedIsbn = allIsbns[0]
// 			}

// 			if strings.Contains(formatedIsbn, ",") {
// 				allIsbns := strings.Split(formatedIsbn, ",")
// 				formatedIsbn = allIsbns[0]
// 			}

// 			/*
// 				Format and insert only ammount and type of size
// 			*/

// 			var formattedAmmount string
// 			var formattedUnit string

// 			size := utils.AnyTypeToString(articlesRetrieved[i]["Size"])

// 			if strings.Contains(size, " ") {
// 				memory := strings.Split(size, " ")
// 				formattedAmmount = memory[0]
// 				formattedUnit = memory[1]
// 			}

// 			AllArticles = append(AllArticles, core.Article{
// 				UniqueID:    utils.AnyTypeToString(articlesRetrieved[i]["UniqueID"]),
// 				SourceName:  utils.AnyTypeToString(articlesRetrieved[i]["SourceName"]),
// 				Url:         utils.AnyTypeToString(articlesRetrieved[i]["Url"]),
// 				DownloadUrl: utils.AnyTypeToString(articlesRetrieved[i]["DownloadUrl"]),
// 				Title:       utils.AnyTypeToString(articlesRetrieved[i]["Title"]),
// 				Search:      utils.AnyTypeToString(articlesRetrieved[i]["Search"]),
// 				Isbn:        formatedIsbn,
// 				Year:        utils.AnyTypeToString(articlesRetrieved[i]["Year"]),
// 				Publisher:   utils.AnyTypeToString(articlesRetrieved[i]["Publisher"]),
// 				Author:      strings.TrimSpace(utils.AnyTypeToString(articlesRetrieved[i]["Author"])),
// 				Extension:   utils.AnyTypeToString(articlesRetrieved[i]["Extension"]),
// 				Page:        utils.AnyTypeToString(articlesRetrieved[i]["Page"]),
// 				Language:    utils.AnyTypeToString(articlesRetrieved[i]["Language"]),
// 				Size:        utils.AnyTypeToString(articlesRetrieved[i]["Size"]),
// 				FileSize: core.Size{
// 					Ammount: formattedAmmount,
// 					Size:    formattedUnit,
// 				},
// 				Time: utils.AnyTypeToString(articlesRetrieved[i]["Time"]),
// 			})
// 		}

// 	fmt.Println(chalk.Green.Color("Articles found by :" + strconv.Itoa(len(AllArticles))))
// 	return AllArticles
// }
