package executions

import (
	"strconv"

	"github.com/Byron/parsecore"
)

func LibrospdfgratismundoExecution() {
	go LibrospdfgratismundoSingleExecution(0, 5)
	go LibrospdfgratismundoSingleExecution(5, 10)
	go LibrospdfgratismundoSingleExecution(10, 15)
	go LibrospdfgratismundoSingleExecution(15, 20)
	go LibrospdfgratismundoSingleExecution(20, 25)
	go LibrospdfgratismundoSingleExecution(25, 30)
	go LibrospdfgratismundoSingleExecution(30, 35)
	go LibrospdfgratismundoSingleExecution(35, 40)
	go LibrospdfgratismundoSingleExecution(40, 45)
	go LibrospdfgratismundoSingleExecution(45, 50)
	go LibrospdfgratismundoSingleExecution(50, 55)
	go LibrospdfgratismundoSingleExecution(55, 60)
	go LibrospdfgratismundoSingleExecution(60, 65)
	go LibrospdfgratismundoSingleExecution(65, 70)
	go LibrospdfgratismundoSingleExecution(70, 75)
	go LibrospdfgratismundoSingleExecution(75, 80)
	go LibrospdfgratismundoSingleExecution(80, 85)
	go LibrospdfgratismundoSingleExecution(85, 90)
	go LibrospdfgratismundoSingleExecution(90, 95)
	go LibrospdfgratismundoSingleExecution(95, 100)
	go LibrospdfgratismundoSingleExecution(100, 105)
	go LibrospdfgratismundoSingleExecution(105, 110)
	go LibrospdfgratismundoSingleExecution(110, 115)
	go LibrospdfgratismundoSingleExecution(115, 120)
	go LibrospdfgratismundoSingleExecution(120, 125)
	go LibrospdfgratismundoSingleExecution(125, 130)
	go LibrospdfgratismundoSingleExecution(130, 135)
	go LibrospdfgratismundoSingleExecution(135, 140)
	go LibrospdfgratismundoSingleExecution(140, 145)
	go LibrospdfgratismundoSingleExecution(145, 150)
	go LibrospdfgratismundoSingleExecution(150, 155)
	go LibrospdfgratismundoSingleExecution(155, 160)
	go LibrospdfgratismundoSingleExecution(160, 165)
	go LibrospdfgratismundoSingleExecution(165, 170)
	go LibrospdfgratismundoSingleExecution(170, 175)
	go LibrospdfgratismundoSingleExecution(175, 180)
	go LibrospdfgratismundoSingleExecution(180, 185)
	go LibrospdfgratismundoSingleExecution(185, 190)
	go LibrospdfgratismundoSingleExecution(190, 195)
	go LibrospdfgratismundoSingleExecution(195, 200)
	go LibrospdfgratismundoSingleExecution(200, 205)
	go LibrospdfgratismundoSingleExecution(205, 210)
	go LibrospdfgratismundoSingleExecution(210, 215)
	go LibrospdfgratismundoSingleExecution(215, 220)
	go LibrospdfgratismundoSingleExecution(220, 225)
	go LibrospdfgratismundoSingleExecution(225, 230)
	go LibrospdfgratismundoSingleExecution(230, 235)
	go LibrospdfgratismundoSingleExecution(235, 240)
	go LibrospdfgratismundoSingleExecution(240, 245)
	go LibrospdfgratismundoSingleExecution(245, 250)
	go LibrospdfgratismundoSingleExecution(250, 255)
	go LibrospdfgratismundoSingleExecution(255, 260)
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
