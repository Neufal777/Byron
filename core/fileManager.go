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

func WriteInFile(name string, Struct interface{}) {

	file, err := json.MarshalIndent(Struct, "", " ")

	if err != nil {
		log.Panic(err)
	}
	_ = ioutil.WriteFile(name+".json", file, 0644)
}

func ReadArticles(category string) []Article {

	// Open our jsonFile
	jsonFile, err := os.Open(category + ".json")

	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
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

		//Create a empty file
		file, err := os.Create("Repository/" + fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		//Write the bytes to the file
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}

	} else {

		fmt.Println(chalk.Magenta.Color("This File already exists! :) "))

	}

	return nil
}
