package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func WriteInFile(File string, Struct interface{}) {
	file, err := json.MarshalIndent(Struct, "", " ")

	if err != nil {
		log.Panic(err)
	}
	_ = ioutil.WriteFile(File, file, 0644)
}

func CheckIfFileExists(ID string) bool {
	info, err := os.Stat("Inventory/" + ID)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
