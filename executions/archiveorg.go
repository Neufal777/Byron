package executions

import (
	"strconv"

	"github.com/Byron/parsecore"
)

func ArchiveOrgExecution() {
	go ArchiveOrgSingleExecution(0, 20)
	go ArchiveOrgSingleExecution(20, 40)
	go ArchiveOrgSingleExecution(40, 60)
	go ArchiveOrgSingleExecution(60, 80)
	go ArchiveOrgSingleExecution(80, 100)
	go ArchiveOrgSingleExecution(100, 120)
	go ArchiveOrgSingleExecution(120, 140)
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
