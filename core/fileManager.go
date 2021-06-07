package core

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ttacon/chalk"
)

func WriteInFile(inventory string, Struct interface{}) {
	file, err := json.MarshalIndent(Struct, "", " ")

	if err != nil {
		log.Panic(err)
	}
	_ = ioutil.WriteFile("Inventory/"+inventory+".json", file, 0644)
}

func ReadArticles(inventory string) []Article {
	jsonFile, err := os.Open("Inventory/" + inventory + ".json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	fileValue, _ := ioutil.ReadAll(jsonFile)

	var Articles []Article

	json.Unmarshal(fileValue, &Articles)
	return Articles
}

func CheckIfFileExists(ID string) bool {
	info, err := os.Stat("Repository/" + ID)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
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
