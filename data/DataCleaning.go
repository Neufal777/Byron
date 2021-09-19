package data

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Byron/core"
	"github.com/Byron/parsecore"
	"github.com/ttacon/chalk"
)

/*
	1- Process all the separated files
	2- build only 1 file
	3- delete ALL duplicates
*/

func DeleteDuplicates(folder, destination string) {
	FreshArticles := map[string]core.Article{}

	var (
		FreshArticlesReady    []core.Article
		duplicates, processed int
		files, _              = ioutil.ReadDir(folder)
		left                  = len(files)
	)

	for _, f := range files {
		fmt.Println(chalk.Green.Color("Processing: " + f.Name()))

		if strings.Contains(f.Name(), ".json") {
			articles := parsecore.ReadArticles(folder + f.Name())

			for _, a := range articles {
				_, ok := FreshArticles[a.Url]
				if !ok && a.Title != "" {
					formatted := a.FormatNewArticle()
					FreshArticles[a.Url] = *formatted
					processed++
				} else {
					duplicates++
				}
			}
		}
		left--
		log.Println("Files Left:", left)
	}

	for _, v := range FreshArticles {
		FreshArticlesReady = append(FreshArticlesReady, v)
	}

	core.WriteInFile(destination, FreshArticlesReady)
	log.Println("Duplicates:", duplicates)
	log.Println("Processed:", processed)
}

func RecoverSource(folder, source string) {
	FreshArticles := map[string]core.Article{}
	/*
		Seperate FILES & ARTICLES by sourcename
	*/

	var (
		FreshArticlesReady []core.Article
		processed          int
	)

	files, _ := ioutil.ReadDir(folder)
	left := len(files)

	for _, f := range files {
		fmt.Println(chalk.Green.Color("Processing: " + f.Name()))

		if strings.Contains(f.Name(), ".json") {
			articles := parsecore.ReadArticles(folder + f.Name())

			for _, a := range articles {
				_, ok := FreshArticles[a.Url]
				if !ok {
					if strings.Contains(a.Url, source) {
						formatted := a.FormatNewArticle()
						FreshArticles[a.Url] = *formatted
						processed++
					}
				}
			}
		}
		left--
		log.Println("Files Left:", left)
	}

	for _, v := range FreshArticles {
		FreshArticlesReady = append(FreshArticlesReady, v)
	}

	core.WriteInFile("UltimateInventory/"+source+"Documents.json", FreshArticlesReady)
	log.Println("Processed:", processed)
}

func GetAllDuplicates(filterby string) {
	files, _ := ioutil.ReadDir("UltimateInventory/")
	var count int

	titleDups := map[string]int{}
	for _, f := range files {

		if strings.Contains(f.Name(), ".json") {

			//1- Open the file
			//2- check dups
			articles := parsecore.ReadArticles("UltimateInventory/" + f.Name())

			for _, a := range articles {
				_, ok := titleDups[a.Title]
				if !ok {
					titleDups[a.Title] = 1
				} else {
					titleDups[a.Title] += 1
				}
			}
		}
	}

	for k, v := range titleDups {
		if v > 10 {
			log.Println(k, ":", v)
			count++
		}
	}

	fmt.Println("Total titles with duplicates:", count)
}
