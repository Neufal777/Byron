package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Byron/core"
	"github.com/ttacon/chalk"
)

func main() {

	Categories := []string{
		"machine learning",
		"english",
		"spanish",
		"artificial Inteligence",
		"neurociencia",
		"inteligencia artificial",
		"math",
		"maths",
		"biology",
		"medicine",
		"clinic",
		"biologia",
		"oxford",
		"harvard",
		"stanford",
		"mit",
		"bio",
		"nano",
		"disease",
		"genomic",
		"vitamine",
		"physics",
		"chemistry",
		"universe",
		"paper",
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(Categories), func(i, j int) { Categories[i], Categories[j] = Categories[j], Categories[i] })

	for _, c := range Categories {

		fmt.Println(chalk.Magenta.Color("processing " + c))
		core.LIBGENDownloadAll(c)

	}

}
