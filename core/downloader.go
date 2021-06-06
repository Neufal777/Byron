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

type Article struct {
	Id          string
	Url         string
	DownloadUrl string
	Title       string
	Isbn        string
	Year        string
	Publisher   string
	Author      string
	Extension   string
	Page        string
	Language    string
	Size        string
	Time        string
}

func LIBGENDownloadAll(search string) {
	var (
		r, _    = regexp.Compile("<a href=.book/index.php.md5=([^\"']*)")
		AllUrls = []string{}
		count   = 0
	)

	for i := 1; i < 102; i++ {

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
		error1, _ := regexp.MatchString(`503 Service Temporarily Unavailable`, htmlFormat)
		error2, _ := regexp.MatchString(`Could not connect to the database`, htmlFormat)

		if !error1 && !error2 {

			matches := r.FindAllStringSubmatch(htmlFormat, -1)

			fmt.Println(chalk.Green.Color("Processing page " + strconv.Itoa(i)))

			if len(matches) > 1 {
				for _, m := range matches {
					fmt.Println(chalk.Green.Color("Saving " + m[1]))
					AllUrls = append(AllUrls, "https://libgen.is/book/index.php?md5="+m[1])
					count++
				}
			} else {
				break
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

	processed := 0

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(AllUrls), func(i, j int) { AllUrls[i], AllUrls[j] = AllUrls[j], AllUrls[i] })

	for _, u := range AllUrls {

		time.Sleep(3 * time.Second)
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

		error1, _ := regexp.MatchString(`503 Service Temporarily Unavailable`, articleHtmlFormat)
		error2, _ := regexp.MatchString(`Could not connect to the database`, articleHtmlFormat)

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

			/*
				Information Display
			*/
			fmt.Println(chalk.White.Color("ID: " + ArticleId.FindStringSubmatch(articleHtmlFormat)[1]))
			fmt.Println(chalk.White.Color("Language:: " + ArticleLang.FindStringSubmatch(articleHtmlFormat)[1]))
			fmt.Println(chalk.White.Color("Download: " + ArticleDownload.FindStringSubmatch(articleHtmlFormat)[1]))
			fmt.Println(chalk.White.Color("Title: " + ArticleTitle.FindStringSubmatch(articleHtmlFormat)[1]))
			fmt.Println(chalk.White.Color("Isbn: " + ArticleIsbn.FindStringSubmatch(articleHtmlFormat)[1]))
			fmt.Println(chalk.White.Color("Year: " + ArticleYear.FindStringSubmatch(articleHtmlFormat)[1]))
			fmt.Println(chalk.White.Color("Publisher: " + ArticlePublisher.FindStringSubmatch(articleHtmlFormat)[1]))
			fmt.Println(chalk.White.Color("Authors: " + ArticleAuthors.FindStringSubmatch(articleHtmlFormat)[1]))
			fmt.Println(chalk.White.Color("Extension: " + ArticleExtension.FindStringSubmatch(articleHtmlFormat)[1]))
			fmt.Println(chalk.White.Color("Pages: " + ArticlePages.FindStringSubmatch(articleHtmlFormat)[1]))
			fmt.Println(chalk.White.Color("Size: " + ArticleSize.FindStringSubmatch(articleHtmlFormat)[1]))
			fmt.Println(chalk.White.Color("Time: " + ArticleTime.FindStringSubmatch(articleHtmlFormat)[1]))

			/*
				Append and download because it's new
			*/
			AllArticles := ReadArticles(search)
			AllArticles = append(AllArticles, Article{
				Id:          ArticleId.FindStringSubmatch(articleHtmlFormat)[1],
				Url:         u,
				DownloadUrl: ArticleDownload.FindStringSubmatch(articleHtmlFormat)[1],
				Title:       ArticleTitle.FindStringSubmatch(articleHtmlFormat)[1],
				Isbn:        ArticleIsbn.FindStringSubmatch(articleHtmlFormat)[1],
				Year:        ArticleYear.FindStringSubmatch(articleHtmlFormat)[1],
				Publisher:   ArticlePublisher.FindStringSubmatch(articleHtmlFormat)[1],
				Author:      ArticleAuthors.FindStringSubmatch(articleHtmlFormat)[1],
				Extension:   ArticleExtension.FindStringSubmatch(articleHtmlFormat)[1],
				Page:        ArticlePages.FindStringSubmatch(articleHtmlFormat)[1],
				Language:    ArticleLang.FindStringSubmatch(articleHtmlFormat)[1],
				Size:        ArticleSize.FindStringSubmatch(articleHtmlFormat)[1],
				Time:        ArticleTime.FindStringSubmatch(articleHtmlFormat)[1],
			})

			fmt.Println(chalk.Green.Color("Added correctly: " + ArticleTitle.FindStringSubmatch(articleHtmlFormat)[1]))
			WriteInFile(search, AllArticles)

			/*
				After that, download the file and save it
			*/
			// FileDownload(
			// 	ArticleDownload.FindStringSubmatch(articleHtmlFormat)[1],  //Download url
			// 	ArticleId.FindStringSubmatch(articleHtmlFormat)[1],        //ID
			// 	ArticleExtension.FindStringSubmatch(articleHtmlFormat)[1], //extension ex. .pdf
			// )
			processed++
			fmt.Println(chalk.Magenta.Color("Processed: " + strconv.Itoa(processed)))
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
