package mongodb

import (
	"fmt"
	"log"
	"strings"

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

	byronDatabase := client.Database("byron")
	byronArticlesCollection := byronDatabase.Collection("byronArticles")

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

	byronDatabase := client.Database("byron")
	byronArticlesCollection := byronDatabase.Collection("byronArticles")

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
	client, ctx, err := ConnectMongoDB()

	var (
		byronDatabase           = client.Database("byron")
		byronArticlesCollection = byronDatabase.Collection("byronArticles")
		AllArticles             []core.Article
	)

	if err != nil {
		log.Fatal(err)
	}

	searchQuery := bson.M{
		"Title": bson.M{"$regex": primitive.Regex{
			Pattern: search,
			Options: "i",
		},
		},
	}

	filterCursor, err := byronArticlesCollection.Find(ctx, searchQuery)

	if err != nil {
		log.Fatal(err)
	}

	var articlesRetrieved []bson.M

	if err = filterCursor.All(ctx, &articlesRetrieved); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(articlesRetrieved); i++ {

		/*
			Format and only show the first Isbn
		*/
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
			Author:      utils.AnyTypeToString(articlesRetrieved[i]["Author"]),
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

	fmt.Println(chalk.Green.Color("Articles found: " + utils.AnyTypeToString(len(AllArticles))))

	return AllArticles
}
