package executions

import (
	"strconv"

	"github.com/Byron/parsecore"
)

func LibrospdfgratismundoExecution() {
	for i := 0; i < 260; i++ {
		//log.Println("Running from:", i, "to", i+5)
		go LibrospdfgratismundoSingleExecution(i, i+5)
		i += 4
	}

	LibrospdfgratismundoSingleExecution(260, 265)
}

func LibrospdfgratismundoSingleExecution(start int, end int) {

	source := parsecore.Source{
		SourceName:           "Librospdfmundo",
		UrlREGEX:             "<h3 class=.titulolibro.><a href=.([^\"']*)",
		TitleREGEX:           "<h1 class=.title.>([^<]*)",
		AuthorREGEX:          "<b>Autores: </b>\\s*<[^>]*>([^<]*)",
		LanguageREGEX:        "<b>Idioma:</b>([^<]*)",
		CompletePageUrlStart: "https://librospdfgratismundo.com/page/",
		ExtensionREGEX:       "<b>Tipo Archivo:</b>([^<]*)",
		AllUrls:              nil,
		Search:               "Librospdfgratismundo:start:" + strconv.Itoa(start) + "end:" + strconv.Itoa(end),
	}

	source.GetArticles(start, end)
}
