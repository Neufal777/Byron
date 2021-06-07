package main

import (
	"flag"
	"fmt"

	"github.com/Byron/core"
	"github.com/ttacon/chalk"
)

func main() {

	// Categories := []string{
	// 	"machine learning",
	// 	"artificial Inteligence",
	// 	"neuroscience",
	// 	"computer",
	// 	"hack",
	// 	"maths",
	// 	"biology",
	// 	"medicine",
	// 	"mit",
	// 	"genomic",
	// 	"physics",
	// 	"fisica",
	// 	"chemistry",
	// 	"universe",
	// 	"paper",
	// }

	category := flag.String("cat", "math", "category")
	flag.Parse()

	// for _, c := range Categories {

	fmt.Println(chalk.Magenta.Color("processing " + *category))
	core.LIBGENDownloadAll(*category)

	// }

}
