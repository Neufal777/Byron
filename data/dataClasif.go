package data

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Byron/core"
	"github.com/ttacon/chalk"
)

/*
	1- Process all the separated files
	2- build only 1 file
	3- delete duplicates
*/

func FilesOrganizer(folder string) {
	ArticlesProcessed := 0
	filesProcessed := 0
	AllArticles := []core.Article{}

	files, _ := ioutil.ReadDir(folder)

	for _, f := range files {
		fmt.Println(chalk.Green.Color("Processing: " + f.Name()))

		if strings.Contains(f.Name(), ".json") {
			file := strings.Replace(f.Name(), ".json", "", -1)
			articles := core.ReadArticles(file)

			for _, art := range articles {
				AllArticles = append(AllArticles, art)
				ArticlesProcessed++
			}
		}
		filesProcessed++
	}

	/*
		Save all the articles in the same file
	*/
	core.WriteInFile("GENERAL", AllArticles)
	log.Println("Total Articles:", ArticlesProcessed)
	log.Println("Total Files:", filesProcessed)

}
