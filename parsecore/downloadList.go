package parsecore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Byron/core"
	"github.com/Byron/utils"
)

func DownloadList(url, search string) {
	allList := ReadDownloads("downloads/" + utils.GetMD5Hash(search) + ".json")
	duplicated := 0

	for _, c := range allList {

		if c == url {
			duplicated = 1
		}
	}

	if duplicated != 1 {
		allList = append(allList, url)
		core.WriteInFile("downloads/"+utils.GetMD5Hash(search)+".json", allList)

	} else {
		log.Println("Duplicated Download")
	}
}

func ReadDownloads(file string) []string {
	var urls []string

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	fileValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(fileValue, &urls)
	return urls
}
