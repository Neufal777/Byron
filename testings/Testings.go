package testings

import (
	"fmt"
	"log"
	"regexp"

	"github.com/Byron/core"
)

func RegexNoResults() {
	/*
		Deal regex with no results
	*/

	html := `<se ction id="block-sitelogo"`
	titleReg, _ := regexp.Compile(`section id=.([^"]*)`)
	title := titleReg.FindStringSubmatch(html)
	var titleFormat string

	if len(title) == 0 {
		log.Println("No results found :(")
	} else {
		titleFormat = titleReg.FindStringSubmatch(html)[1]
		log.Println("Title:", titleFormat)
	}

}

func SingleLoad(filename string) {

	file := core.CheckIfFileExists(filename)

	if file {
		fmt.Println("this file already exists!")
	}

	log.Println("Done!")
}
