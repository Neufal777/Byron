package executions

import "github.com/Byron/parsecore"

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

	go ManyBooksSingleExecution(categories[0])
	go ManyBooksSingleExecution(categories[1])
	go ManyBooksSingleExecution(categories[2])
	go ManyBooksSingleExecution(categories[3])
	go ManyBooksSingleExecution(categories[4])
	go ManyBooksSingleExecution(categories[5])
	go ManyBooksSingleExecution(categories[6])
	go ManyBooksSingleExecution(categories[7])
	go ManyBooksSingleExecution(categories[8])
	go ManyBooksSingleExecution(categories[9])
	go ManyBooksSingleExecution(categories[10])
	go ManyBooksSingleExecution(categories[11])
	go ManyBooksSingleExecution(categories[12])
	go ManyBooksSingleExecution(categories[13])
	go ManyBooksSingleExecution(categories[14])
	go ManyBooksSingleExecution(categories[15])
	go ManyBooksSingleExecution(categories[16])
	go ManyBooksSingleExecution(categories[17])
	go ManyBooksSingleExecution(categories[18])
	go ManyBooksSingleExecution(categories[19])
	go ManyBooksSingleExecution(categories[20])
	go ManyBooksSingleExecution(categories[21])
	go ManyBooksSingleExecution(categories[22])
	go ManyBooksSingleExecution(categories[23])
	go ManyBooksSingleExecution(categories[24])
	go ManyBooksSingleExecution(categories[25])
	go ManyBooksSingleExecution(categories[26])
	go ManyBooksSingleExecution(categories[27])
	go ManyBooksSingleExecution(categories[28])
	go ManyBooksSingleExecution(categories[29])
	go ManyBooksSingleExecution(categories[30])
	go ManyBooksSingleExecution(categories[31])
	go ManyBooksSingleExecution(categories[32])
	go ManyBooksSingleExecution(categories[33])
	go ManyBooksSingleExecution(categories[34])
	go ManyBooksSingleExecution(categories[35])
	go ManyBooksSingleExecution(categories[36])
	go ManyBooksSingleExecution(categories[37])
	go ManyBooksSingleExecution(categories[38])
	go ManyBooksSingleExecution(categories[39])
	go ManyBooksSingleExecution(categories[40])
	go ManyBooksSingleExecution(categories[41])
	go ManyBooksSingleExecution(categories[42])
	go ManyBooksSingleExecution(categories[43])
	go ManyBooksSingleExecution(categories[44])
	go ManyBooksSingleExecution(categories[45])
	go ManyBooksSingleExecution(categories[46])
	go ManyBooksSingleExecution(categories[47])
	go ManyBooksSingleExecution(categories[48])
	go ManyBooksSingleExecution(categories[49])
	go ManyBooksSingleExecution(categories[50])
	go ManyBooksSingleExecution(categories[51])
	go ManyBooksSingleExecution(categories[52])
	go ManyBooksSingleExecution(categories[53])
	go ManyBooksSingleExecution(categories[54])
	go ManyBooksSingleExecution(categories[55])
	go ManyBooksSingleExecution(categories[56])
	go ManyBooksSingleExecution(categories[57])
	go ManyBooksSingleExecution(categories[58])
	go ManyBooksSingleExecution(categories[59])
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
