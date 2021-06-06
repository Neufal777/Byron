package main

import (
	"fmt"

	"github.com/Byron/core"
	"github.com/ttacon/chalk"
)

func main() {

	Categories := []string{
		//"machine learning",
		//"artificial Inteligence",
		//"neuroscience",
		//"maths",
		//"biology",
		//"medicine",
		//"mit",
		"genomic",
		//"physics",
		//"chemistry",
		//"universe",
		//"paper",
	}

	// rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(len(Categories), func(i, j int) { Categories[i], Categories[j] = Categories[j], Categories[i] })

	for _, c := range Categories {

		fmt.Println(chalk.Magenta.Color("processing " + c))
		core.LIBGENDownloadAll(c)

	}

}
