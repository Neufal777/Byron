package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/Byron/core"
	"github.com/Byron/utils"
	"github.com/ttacon/chalk"
)

const PAGE_LIMIT = 4

type Source struct {
	SourceName           string
	UrlREGEX             string
	IdREGEX              string
	SearchREGEX          string
	DownloadUrlREGEX     string
	DownloadUrlComplete  string
	TitleREGEX           string
	IsbnREGEX            string
	YearREGEX            string
	PublisherREGEX       string
	AuthorREGEX          string
	ExtensionREGEX       string
	PageREGEX            string
	LanguageREGEX        string
	SizeREGEX            string
	TimeREGEX            string
	CompletePageUrlStart string
	CompletePageUrlEnd   string
	IncompleteArticleUrl string
	AllUrls              []string
	Search               string
}

func (s *Source) GetArticles() {
	r, _ := regexp.Compile(s.UrlREGEX)
	processed := 0

	for i := 1; i < PAGE_LIMIT; i++ {
		time.Sleep(1 * time.Second)

		htmlFormat, errs := ProxyScraping(s.CompletePageUrlStart + strconv.Itoa(i) + s.CompletePageUrlEnd)

		if errs != nil {
			log.Println("Ups, we have some errors")

		}
		if !core.ErrorsHandling(htmlFormat) {
			matches := r.FindAllStringSubmatch(htmlFormat, -1)
			fmt.Println(chalk.Green.Color("Processing page " + strconv.Itoa(i)))

			if len(matches) < 1 {
				break
			}

			for _, m := range matches {
				fmt.Println(chalk.Green.Color("Saving " + m[1]))
				s.AllUrls = append(s.AllUrls, s.IncompleteArticleUrl+m[1])
				processed++

				//DownloadList(s.IncompleteArticleUrl+m[1], s.Search)
			}

		} else {
			fmt.Println(chalk.Magenta.Color("Given 503. waiting to reconnect"))
			time.Sleep(10 * time.Second)
		}
	}
	s.ProcessArticles()
}

func (s *Source) ProcessArticles() {
	fmt.Println(chalk.Green.Color("Start processing Articles.."))

	log.Println("Num of Total Articles:", len(s.AllUrls))
	processed := 0

	//randomize urls processing
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s.AllUrls), func(i, j int) { s.AllUrls[i], s.AllUrls[j] = s.AllUrls[j], s.AllUrls[i] })

	for _, u := range s.AllUrls {
		time.Sleep(1 * time.Second)

		articleHtmlFormat, errs := ProxyScraping(u)
		if errs != nil {
			log.Println("Ups, we have some errors")

		}

		if !core.ErrorsHandling(articleHtmlFormat) {

			newArticle := core.Article{
				SourceName: s.SourceName,
				Url:        u,
				Search:     s.Search,
			}

			log.Println("Article:", u)
			newArticle = CheckRegex(s, newArticle, articleHtmlFormat)

			/*
				Append and download because it's new
			*/

			AllArticles := ReadArticles("Inventory/" + utils.GetMD5Hash(s.Search) + ".json")
			newArticleFormatted := newArticle.FormatNewArticle()

			dup := CountDuplicates(AllArticles, newArticleFormatted)

			if dup == 0 {
				AllArticlesUpdated := ReadArticles("Inventory/" + utils.GetMD5Hash(s.Search) + ".json")
				AllArticlesUpdated = append(AllArticlesUpdated, *newArticleFormatted)
				core.WriteInFile("Inventory/"+utils.GetMD5Hash(s.Search)+".json", AllArticlesUpdated)

				/*
					Display relevant information about the new Document
				*/
				newArticle.DisplayInformation()
				fmt.Println(chalk.Green.Color("Added correctly: " + newArticle.Title))
				processed++
				fmt.Println(chalk.Magenta.Color("Processed: " + strconv.Itoa(processed)))

				/*
					After that, download the file and save it
				*/
				// core.FileDownload(
				// 	ArticleDownload.FindStringSubmatch(articleHtmlFormat)[1],  //Download url
				// 	ArticleId.FindStringSubmatch(articleHtmlFormat)[1],        //ID
				// 	ArticleExtension.FindStringSubmatch(articleHtmlFormat)[1], //extension ex. .pdf
				// )

			} else {
				fmt.Println(chalk.Red.Color("This article already exists, nothing to do here"))
			}

		} else {
			fmt.Println(chalk.Magenta.Color("Given 503. waiting to reconnect"))

		}
	}

	fmt.Println(chalk.Green.Color("All the documents were Downloaded :) "))
}

func CountDuplicates(AllArticles []core.Article, newArticleFormatted *core.Article) int {
	duplicated := 0
	for i := 0; i < len(AllArticles); i++ {
		if AllArticles[i].Url == newArticleFormatted.Url {
			duplicated = 1
			break
		}
	}
	return duplicated
}

func ReadArticles(file string) []core.Article {
	var Articles []core.Article
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	fileValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(fileValue, &Articles)
	return Articles
}

func CheckRegex(s *Source, newArticle core.Article, articleHtmlFormat string) core.Article {
	if regexSet(s.TitleREGEX) {
		ArticleTitle, _ := regexp.Compile(s.TitleREGEX)
		newArticle.Title = ArticleTitle.FindStringSubmatch(articleHtmlFormat)[1]
	}
	if regexSet(s.AuthorREGEX) {
		ArticleAuthors, _ := regexp.Compile(s.AuthorREGEX)
		newArticle.Author = ArticleAuthors.FindStringSubmatch(articleHtmlFormat)[1]
	}
	if regexSet(s.PublisherREGEX) {
		ArticlePublisher, _ := regexp.Compile(s.PublisherREGEX)
		newArticle.Publisher = ArticlePublisher.FindStringSubmatch(articleHtmlFormat)[1]
	}
	if regexSet(s.YearREGEX) {
		ArticleYear, _ := regexp.Compile(s.YearREGEX)
		newArticle.Year = ArticleYear.FindStringSubmatch(articleHtmlFormat)[1]
	}
	if regexSet(s.LanguageREGEX) {
		ArticleLang, _ := regexp.Compile(s.LanguageREGEX)
		newArticle.Language = ArticleLang.FindStringSubmatch(articleHtmlFormat)[1]
	}
	if regexSet(s.IsbnREGEX) {
		ArticleIsbn, _ := regexp.Compile(s.IsbnREGEX)
		newArticle.Isbn = ArticleIsbn.FindStringSubmatch(articleHtmlFormat)[1]
	}
	if regexSet(s.TimeREGEX) {
		ArticleTime, _ := regexp.Compile(s.TimeREGEX)
		newArticle.Time = ArticleTime.FindStringSubmatch(articleHtmlFormat)[1]
	}
	if regexSet(s.IdREGEX) {
		ArticleId, _ := regexp.Compile(s.IdREGEX)
		newArticle.Id = ArticleId.FindStringSubmatch(articleHtmlFormat)[1]
	}
	if regexSet(s.SizeREGEX) {
		ArticleSize, _ := regexp.Compile(s.SizeREGEX)
		newArticle.Size = ArticleSize.FindStringSubmatch(articleHtmlFormat)[1]
	}
	if regexSet(s.PageREGEX) {
		ArticlePages, _ := regexp.Compile(s.PageREGEX)
		newArticle.Page = ArticlePages.FindStringSubmatch(articleHtmlFormat)[1]
	}
	if regexSet(s.ExtensionREGEX) {
		ArticleExtension, _ := regexp.Compile(s.ExtensionREGEX)
		newArticle.Extension = ArticleExtension.FindStringSubmatch(articleHtmlFormat)[1]
	}
	if regexSet(s.DownloadUrlREGEX) {
		ArticleDownload, _ := regexp.Compile(s.DownloadUrlREGEX)
		newArticle.DownloadUrl = s.DownloadUrlComplete + ArticleDownload.FindStringSubmatch(articleHtmlFormat)[1]

	}

	return newArticle
}

func regexSet(regex string) bool {
	return regex != ""
}
