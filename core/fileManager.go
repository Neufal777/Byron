package core

import (
	"encoding/json"
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

func CheckIfFileExists(ID string) bool {
	info, err := os.Stat("Repository/" + ID)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
