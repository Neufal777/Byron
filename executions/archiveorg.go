package executions

import (
	"strconv"

	"github.com/Byron/parsecore"
)

func ArchiveOrgExecution() {
	for i := 0; i < 140; i++ {
		//log.Println("Running from:", i, "to", i+20)
		go ArchiveOrgSingleExecution(i, i+20)
		i += 19
	}

	ArchiveOrgSingleExecution(140, 180)
}

func ArchiveOrgSingleExecution(start int, end int) {

	source := parsecore.Source{
		SourceName:           "ArchiveOrg",
		UrlREGEX:             "<a href=./details/([^\"']*)",
		IncompleteArticleUrl: "https://archive.org/details/",
		TitleREGEX:           "<span class=.breaker-breaker. itemprop=.name.>([^<]*)",
		YearREGEX:            "<span itemprop=.datePublished.>([^<]*)",
		PublisherREGEX:       "<span itemprop=.publisher.>([^<]*)",
		AuthorREGEX:          "/search.php.query=creator[^>]*>([^<]*)",
		CompletePageUrlStart: "https://archive.org/details/books?and%5B%5D=languageSorter%3A%22English%22&and%5B%5D=languageSorter%3A%22Spanish%22&and%5B%5D=languageSorter%3A%22Italian%22&sort=-week&page=",
		AllUrls:              nil,
		Search:               "Start:" + strconv.Itoa(start) + "ArchiveOrg" + strconv.Itoa(end),
	}

	source.GetArticles(start, end)
}
