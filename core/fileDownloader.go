package core

/*
	Source : Libgen.is
*/
import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/ttacon/chalk"
)

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

	log.Println(DownloadHtmlFormat)
	if !ErrorsHandling(DownloadHtmlFormat) {
		downloadlinkRegex, _ := regexp.Compile("<h2><a href=.([^\"']*)")
		downloadlink := downloadlinkRegex.FindStringSubmatch(DownloadHtmlFormat)[1]

		DownloadPDF(downloadlink, ID+"."+format)
	} else {
		fmt.Println(chalk.Magenta.Color("Given connection error, waiting to reconnect"))
		time.Sleep(5 * time.Second)
	}

}

func DownloadPDF(URL, fileName string) error {
	exists := CheckIfFileExists(fileName)

	if !exists {

		fmt.Println(chalk.Magenta.Color("Downloading File.. "))
		response, err := http.Get(URL)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		file, err := os.Create("Repository/" + fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}

	} else {
		fmt.Println(chalk.Magenta.Color("This File already exists! :) "))
	}

	return nil
}
