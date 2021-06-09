package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
