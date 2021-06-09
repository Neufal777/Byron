package main

import "github.com/Byron/data"

func main() {

	// category := flag.String("cat", "math", "category")
	// flag.Parse()
	// fmt.Println(chalk.Magenta.Color("processing " + *category))

	// excat := []string{"salsas", "botas"}

	// source := sources.Source{
	// 	SourceName:           "LIBGEN",
	// 	UrlREGEX:             "<a href=.book/index.php.md5=([^\"']*)",
	// 	IdREGEX:              "ID:</font></nobr></td><td>([^<]*)",
	// 	DownloadUrlREGEX:     "align=.center.><a href=.([^\"']*). title=.Gen.lib.rus.ec.",
	// 	TitleREGEX:           "<title>Library Genesis:([^<]*)",
	// 	IsbnREGEX:            "ISBN:</font></td><td>([^<]*)",
	// 	YearREGEX:            "Year:</font></nobr></td><td>([^<]*)",
	// 	PublisherREGEX:       "Publisher:</font></nobr></td><td>([^<]*)",
	// 	AuthorREGEX:          "Author.s.:</font></nobr></td><td colspan=3><b>([^<]*)",
	// 	ExtensionREGEX:       "Extension:</font></nobr></td><td>([^<]*)",
	// 	PageREGEX:            "Pages .biblio.tech.:</font></nobr></td><td>([^<]*)",
	// 	LanguageREGEX:        "Language:</font></nobr></td><td>([^<]*)",
	// 	SizeREGEX:            "Size:</font></nobr></td><td>([^<]*)",
	// 	TimeREGEX:            "Time modified:</font></nobr></td><td>([^<]*)",
	// 	CompletePageUrl:      "https://libgen.is/search.php?&res=100&req=" + excat[0] + "&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
	// 	IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
	// 	AllUrls:              nil,
	// 	Search:               excat[0],
	// }

	// source.GetArticles()

	data.FilesOrganizer("Inventory/")
}
