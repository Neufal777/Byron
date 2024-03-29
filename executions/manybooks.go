package executions

import (
	"log"

	"github.com/Byron/parsecore"
)

func ManyBooksExecution() {
	categories := []string{
		"https://manybooks.net/categories/ADV",
		"https://manybooks.net/categories/AFR",
		"https://manybooks.net/categories/ART",
		"https://manybooks.net/categories/BAN",
		"https://manybooks.net/categories/BIO",
		"https://manybooks.net/categories/BUS",
		"https://manybooks.net/categories/CAN",
		"https://manybooks.net/categories/CLA",
		"https://manybooks.net/categories/COM",
		"https://manybooks.net/categories/COO",
		"https://manybooks.net/categories/COR",
		"https://manybooks.net/categories/CCL",
		"https://manybooks.net/categories/CRI",
		"https://manybooks.net/categories/DRA",
		"https://manybooks.net/categories/SPY",
		"https://manybooks.net/categories/ESS",
		"https://manybooks.net/categories/ETT",
		"https://manybooks.net/categories/FAN",
		"https://manybooks.net/categories/FIC",
		"https://manybooks.net/categories/GAM",
		"https://manybooks.net/categories/GAY",
		"https://manybooks.net/categories/GHO",
		"https://manybooks.net/categories/GOT",
		"https://manybooks.net/categories/GOV",
		"https://manybooks.net/categories/HAR",
		"https://manybooks.net/categories/HEA",
		"https://manybooks.net/categories/HIS",
		"https://manybooks.net/categories/HOR",
		"https://manybooks.net/categories/HUM",
		"https://manybooks.net/categories/INS",
		"https://manybooks.net/categories/LAN",
		"https://manybooks.net/categories/MUS",
		"https://manybooks.net/categories/MYS",
		"https://manybooks.net/categories/MYT",
		"https://manybooks.net/categories/NAT",
		"https://manybooks.net/categories/NAU",
		"https://manybooks.net/categories/NON",
		"https://manybooks.net/categories/OCC",
		"https://manybooks.net/categories/PER",
		"https://manybooks.net/categories/PHI",
		"https://manybooks.net/categories/PIR",
		"https://manybooks.net/categories/POE",
		"https://manybooks.net/categories/POL",
		"https://manybooks.net/categories/MOD",
		"https://manybooks.net/categories/PSY",
		"https://manybooks.net/categories/PUL",
		"https://manybooks.net/categories/RAN",
		"https://manybooks.net/categories/REF",
		"https://manybooks.net/categories/REL",
		"https://manybooks.net/categories/ROM",
		"https://manybooks.net/categories/SAT",
		"https://manybooks.net/categories/SCI",
		"https://manybooks.net/categories/SFC",
		"https://manybooks.net/categories/SST",
		"https://manybooks.net/categories/SHO",
		"https://manybooks.net/categories/THR",
		"https://manybooks.net/categories/TRA",
		"https://manybooks.net/categories/WAR",
		"https://manybooks.net/categories/WES",
		"https://manybooks.net/categories/WOM",
		"https://manybooks.net/categories/CHI",
	}

	for i := 0; i < len(categories)-1; i++ {
		log.Println("Running: ", categories[i], "ID num:", i)
		go ManyBooksSingleExecution(categories[i])

	}

	ManyBooksSingleExecution(categories[60])
}

func ManyBooksSingleExecution(urlCat string) {
	source := parsecore.Source{
		SourceName:           "ManyBooks",
		UrlREGEX:             "about=./titles/([^\"']*)",
		IncompleteArticleUrl: "https://manybooks.net/titles/",
		DownloadUrlREGEX:     "<div class=.form-group.><a href=./book/([^\"']*)",
		DownloadUrlComplete:  "https://manybooks.net/book/",
		TitleREGEX:           "<h1 class=.page-header.>([^<]*)",
		YearREGEX:            "<div class=.field field--name-field-published-year field--type-integer field--label-hidden field--item.>([^<]*)",
		AuthorREGEX:          "itemprop=.author.[^>]*>([^<]*)",
		PageREGEX:            "<div class=.field field--name-field-pages field--type-integer field--label-hidden field--item.>([^<]*)",
		LanguageREGEX:        "language: .([^\"']*)",
		CompletePageUrlStart: urlCat + "?language=All&sort_by=field_downloads&page=",
		CompletePageUrlEnd:   "",
		AllUrls:              nil,
		Search:               urlCat,
	}

	source.GetArticles(1, 400)
}
