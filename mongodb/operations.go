package mongodb

import (
	"fmt"
	"log"

	"github.com/Byron/core"
	"github.com/Byron/utils"
	"github.com/ttacon/chalk"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertArticle(article *core.Article) error {

	client, ctx, err := ConnectMongoDB()

	defer client.Disconnect(ctx)

	if err != nil {
		panic(err)
	}

	byronDatabase := client.Database("byron")
	byronArticlesCollection := byronDatabase.Collection("byronArticles")

	_, err = byronArticlesCollection.InsertOne(ctx, bson.D{
		{Key: "uniqueid", Value: article.UniqueID},
		{Key: "sourcename", Value: article.SourceName},
		{Key: "url", Value: article.Url},
		{Key: "downloadurl", Value: article.DownloadUrl},
		{Key: "title", Value: article.Title},
		{Key: "search", Value: article.Search},
		{Key: "isbn", Value: article.Isbn},
		{Key: "year", Value: article.Year},
		{Key: "publisher", Value: article.Publisher},
		{Key: "author", Value: article.Author},
		{Key: "extension", Value: article.Extension},
		{Key: "page", Value: article.Page},
		{Key: "language", Value: article.Language},
		{Key: "size", Value: article.Size},
		{Key: "time", Value: article.Time},
	})

	if err != nil {
		log.Println(err)
	}

	fmt.Println(chalk.Green.Color("Inserted correctly: " + article.UniqueID))
	return err

}

func RetrieveArticle(url string) []core.Article {
	var AllArticles []core.Article
	client, ctx, err := ConnectMongoDB()

	if err != nil {
		log.Fatal(err)
		log.Println("CONEXION ERROOOOR!")
	}

	byronDatabase := client.Database("byron")
	byronArticlesCollection := byronDatabase.Collection("byronArticles")

	filterCursor, err := byronArticlesCollection.Find(ctx, bson.M{"sourcename": url})
	if err != nil {
		log.Fatal(err)
	}
	var articlesRetrieved []bson.M
	if err = filterCursor.All(ctx, &articlesRetrieved); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(articlesRetrieved); i++ {
		AllArticles = append(AllArticles, core.Article{
			UniqueID:    utils.AnyTypeToString(articlesRetrieved[i]["uniqueid"]),
			SourceName:  utils.AnyTypeToString(articlesRetrieved[i]["sourcename"]),
			Url:         utils.AnyTypeToString(articlesRetrieved[i]["url"]),
			DownloadUrl: utils.AnyTypeToString(articlesRetrieved[i]["downloadurl"]),
			Title:       utils.AnyTypeToString(articlesRetrieved[i]["title"]),
			Search:      utils.AnyTypeToString(articlesRetrieved[i]["search"]),
			Isbn:        utils.AnyTypeToString(articlesRetrieved[i]["isbn"]),
			Year:        utils.AnyTypeToString(articlesRetrieved[i]["year"]),
			Publisher:   utils.AnyTypeToString(articlesRetrieved[i]["publisher"]),
			Author:      utils.AnyTypeToString(articlesRetrieved[i]["author"]),
			Extension:   utils.AnyTypeToString(articlesRetrieved[i]["extension"]),
			Page:        utils.AnyTypeToString(articlesRetrieved[i]["page"]),
			Language:    utils.AnyTypeToString(articlesRetrieved[i]["language"]),
			Size:        utils.AnyTypeToString(articlesRetrieved[i]["size"]),
			Time:        utils.AnyTypeToString(articlesRetrieved[i]["time"]),
		})
	}
	return AllArticles
}
