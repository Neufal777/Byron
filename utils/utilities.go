package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
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

func AnyTypeToString(input interface{}) string {
	return fmt.Sprintf("%v", input)
}

func FixUnitedNames(name string) string {
	var formatted string
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)

	submatchall := re.FindAllString(name, -1)
	for _, element := range submatchall {
		formatted += element + " "
	}

	/*
		Formatting final name
	*/
	formatted = strings.ReplaceAll(formatted, "-", "")
	formatted = strings.ReplaceAll(formatted, " )", ")")
	formatted = strings.ReplaceAll(formatted, "(", " (")
	formatted = strings.ReplaceAll(formatted, "  ", " ")
	formatted = strings.TrimSpace(formatted)

	return formatted
}
