package sources

func AllSourcesScrapingInformation() {

	_ = Source{
		SourceName:           "OpenLibra",
		UrlREGEX:             "<div class=.image-cover.> <a title=.[^<]*href=.([^\"']*)",
		IdREGEX:              "",
		DownloadUrlREGEX:     "<a class=.command-button btn btn-ol-twitter. href=.([^\"']*)",
		TitleREGEX:           "<h1 itemprop=.name. class[^>]*>([^<]*)",
		IsbnREGEX:            "",
		YearREGEX:            "<td itemprop=.copyrightYear.>([^<]*)",
		PublisherREGEX:       "<td itemprop=.publisher[^<]*<span itemprop=.name.>([^<]*)",
		AuthorREGEX:          "<span itemprop=.author.>([^<]*)",
		ExtensionREGEX:       "",
		PageREGEX:            "<td itemprop=.numberOfPages.>([^<]*)",
		LanguageREGEX:        "<td>Idioma:</td><td>([^<]*)",
		SizeREGEX:            "<td>Tama.o:</td><td>([^<]*)",
		TimeREGEX:            "",
		CompletePageUrl:      "https://openlibra.com/es/collection/pag/",
		IncompleteArticleUrl: "",
		AllUrls:              nil,
		Search:               "OpenLibraGeneral",
	}

	_ = Source{
		SourceName:           "LIBGEN",
		UrlREGEX:             "<a href=.book/index.php.md5=([^\"']*)",
		IdREGEX:              "ID:</font></nobr></td><td>([^<]*)",
		DownloadUrlREGEX:     "align=.center.><a href=.([^\"']*). title=.Gen.lib.rus.ec.",
		TitleREGEX:           "<title>Library Genesis:([^<]*)",
		IsbnREGEX:            "ISBN:</font></td><td>([^<]*)",
		YearREGEX:            "Year:</font></nobr></td><td>([^<]*)",
		PublisherREGEX:       "Publisher:</font></nobr></td><td>([^<]*)",
		AuthorREGEX:          "Author.s.:</font></nobr></td><td colspan=3><b>([^<]*)",
		ExtensionREGEX:       "Extension:</font></nobr></td><td>([^<]*)",
		PageREGEX:            "Pages .biblio.tech.:</font></nobr></td><td>([^<]*)",
		LanguageREGEX:        "Language:</font></nobr></td><td>([^<]*)",
		SizeREGEX:            "Size:</font></nobr></td><td>([^<]*)",
		TimeREGEX:            "Time modified:</font></nobr></td><td>([^<]*)",
		CompletePageUrl:      "https://libgen.is/search.php?&res=100&req=QUERY&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
		IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
		AllUrls:              nil,
		Search:               "machine learning",
	}
}
