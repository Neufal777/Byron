package core

/*
	Source : Libgen.is
*/
import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/ttacon/chalk"
)

type Size struct {
	Ammount string
	Size    string
}
type Article struct {
	Id           string
	UniqueID     string
	Url          string
	Search       string
	DownloadUrl  string
	Title        string
	Isbn         string
	IsbnList     []string
	Year         string
	Publisher    string
	Author       string
	AuthorList   []string
	Extension    string
	Page         string
	Language     string
	LanguageList []string
	Size         string
	FileSize     Size
	Time         string
}

func LIBGENDownloadAll(search string) {

	var (
		r, _             = regexp.Compile("<a href=.book/index.php.md5=([^\"']*)")
		error1Compile, _ = regexp.Compile("503 Service Temporarily Unavailable")
		error2Compile, _ = regexp.Compile("Could not connect to the database")
		AllUrls          = []string{}
		count            = 0
	)

	for i := 1; i < 103; i++ {

		resp, err := http.Get("https://libgen.is/search.php?&res=100&req=" + search + "&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=" + strconv.Itoa(i))
		//resp, err := http.Get("https://libgen.is/search.php?mode=last&view=simple&phrase=0&timefirst=&timelast=&sort=def&sortmode=ASC&page=" + strconv.Itoa(i))

		if err != nil {
			log.Println(err)
		}

		defer resp.Body.Close()
		html, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		htmlFormat := string(html)

		/*
			Check if the website return any kind of error
		*/

		error1 := error1Compile.MatchString(htmlFormat)
		error2 := error2Compile.MatchString(htmlFormat)

		if !error1 && !error2 {

			matches := r.FindAllStringSubmatch(htmlFormat, -1)
			fmt.Println(chalk.Green.Color("Processing page " + strconv.Itoa(i)))

			for _, m := range matches {

				fmt.Println(chalk.Green.Color("Saving " + m[1]))
				AllUrls = append(AllUrls, "https://libgen.is/book/index.php?md5="+m[1])
				count++
			}

		} else {
			fmt.Println(chalk.Magenta.Color("Given 503. waiting to reconnect"))
			time.Sleep(10 * time.Second)
		}

	}

	log.Println("Total Articles:", count)
	ProcessUrls(AllUrls, search)
}

func ProcessUrls(AllUrls []string, search string) {

	var (
		processed        = 0
		error1Compile, _ = regexp.Compile("503 Service Temporarily Unavailable")
		error2Compile, _ = regexp.Compile("Could not connect to the database")
	)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(AllUrls), func(i, j int) { AllUrls[i], AllUrls[j] = AllUrls[j], AllUrls[i] })

	for _, u := range AllUrls {
		time.Sleep(2 * time.Second)
		resp, err := http.Get(u)

		if err != nil {
			log.Println(err)
		}

		defer resp.Body.Close()

		articleHtml, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		articleHtmlFormat := string(articleHtml)

		/*
			Check for any error in the website
		*/
		error1 := error1Compile.MatchString(articleHtmlFormat)
		error2 := error2Compile.MatchString(articleHtmlFormat)

		if !error1 && !error2 {

			ArticleTitle, _ := regexp.Compile("<title>Library Genesis:([^<]*)")
			ArticleAuthors, _ := regexp.Compile("Author.s.:</font></nobr></td><td colspan=3><b>([^<]*)")
			ArticlePublisher, _ := regexp.Compile("Publisher:</font></nobr></td><td>([^<]*)")
			ArticleYear, _ := regexp.Compile("Year:</font></nobr></td><td>([^<]*)")
			ArticleLang, _ := regexp.Compile("Language:</font></nobr></td><td>([^<]*)")
			ArticleIsbn, _ := regexp.Compile("ISBN:</font></td><td>([^<]*)")
			ArticleTime, _ := regexp.Compile("Time modified:</font></nobr></td><td>([^<]*)")
			ArticleSize, _ := regexp.Compile("Size:</font></nobr></td><td>([^<]*)")
			ArticlePages, _ := regexp.Compile("Pages .biblio.tech.:</font></nobr></td><td>([^<]*)")
			ArticleId, _ := regexp.Compile("ID:</font></nobr></td><td>([^<]*)")
			ArticleExtension, _ := regexp.Compile("Extension:</font></nobr></td><td>([^<]*)")
			ArticleDownload, _ := regexp.Compile("align=.center.><a href=.([^\"']*). title=.Gen.lib.rus.ec.")

			newArticle := Article{
				Id:          ArticleId.FindStringSubmatch(articleHtmlFormat)[1],
				Url:         u,
				Search:      search,
				UniqueID:    "",
				DownloadUrl: ArticleDownload.FindStringSubmatch(articleHtmlFormat)[1],
				Title:       ArticleTitle.FindStringSubmatch(articleHtmlFormat)[1],
				Isbn:        ArticleIsbn.FindStringSubmatch(articleHtmlFormat)[1],
				IsbnList:    nil,
				Year:        ArticleYear.FindStringSubmatch(articleHtmlFormat)[1],
				Publisher:   ArticlePublisher.FindStringSubmatch(articleHtmlFormat)[1],
				Author:      ArticleAuthors.FindStringSubmatch(articleHtmlFormat)[1],
				Extension:   ArticleExtension.FindStringSubmatch(articleHtmlFormat)[1],
				Page:        ArticlePages.FindStringSubmatch(articleHtmlFormat)[1],
				Language:    ArticleLang.FindStringSubmatch(articleHtmlFormat)[1],
				Size:        ArticleSize.FindStringSubmatch(articleHtmlFormat)[1],
				Time:        ArticleTime.FindStringSubmatch(articleHtmlFormat)[1],
			}

			/*
				Append and download because it's new
			*/

			AllArticles := ReadArticles(FILENAME)

			newArticleFormatted := newArticle.FormatNewArticle()

			duplicated := 0
			for i := 0; i < len(AllArticles); i++ {

				if AllArticles[i].Id == newArticleFormatted.Id {
					duplicated = 1
					break
				}
			}

			if duplicated == 0 {

				AllArticlesUpdated := ReadArticles(FILENAME)
				AllArticlesUpdated = append(AllArticlesUpdated, *newArticleFormatted)
				WriteInFile(FILENAME, AllArticlesUpdated)

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
				// FileDownload(
				// 	ArticleDownload.FindStringSubmatch(articleHtmlFormat)[1],  //Download url
				// 	ArticleId.FindStringSubmatch(articleHtmlFormat)[1],        //ID
				// 	ArticleExtension.FindStringSubmatch(articleHtmlFormat)[1], //extension ex. .pdf
				// )

			} else {
				fmt.Println(chalk.Red.Color("This article already exists, nothing to do here"))
			}

		} else {

			fmt.Println(chalk.Magenta.Color("Given 503. waiting to reconnect"))
			time.Sleep(10 * time.Second)
		}

	}

	fmt.Println(chalk.Green.Color("All the documents were Downloaded :) "))
}

func FileDownload(URL, ID, format string) {

	/*
		Check html, regex the download link and save the file :)
	*/

	resp, err := http.Get(URL)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	DownloadHtml, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	DownloadHtmlFormat := string(DownloadHtml)
	error1, _ := regexp.MatchString(`503 Service Temporarily Unavailable`, DownloadHtmlFormat)
	error2, _ := regexp.MatchString(`Could not connect to the database`, DownloadHtmlFormat)

	if !error1 && !error2 {
		downloadlinkRegex, _ := regexp.Compile("<h2><a href=.([^\"']*)")
		downloadlink := downloadlinkRegex.FindStringSubmatch(DownloadHtmlFormat)[1]

		DownloadPDF(downloadlink, ID+"."+format)
	} else {
		fmt.Println(chalk.Magenta.Color("Given connection error, waiting to reconnect"))
		time.Sleep(10 * time.Second)
	}

}
