package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func PrettyPrintStruct(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func Render(w http.ResponseWriter, filename string, data interface{}) {

	tmpl, err := template.ParseFiles(filename)

	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}
