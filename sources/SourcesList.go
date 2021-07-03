package sources

import "github.com/Byron/parsecore"

func AllSourcesScrapingInformation() {

	_ = parsecore.Source{
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
		CompletePageUrlStart: "https://openlibra.com/es/collection/pag/",
		IncompleteArticleUrl: "",
		AllUrls:              nil,
		Search:               "OpenLibraGeneral",
	}

	_ = parsecore.Source{
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
		CompletePageUrlStart: "https://libgen.is/search.php?&res=100&req=QUERY&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
		IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
		AllUrls:              nil,
		Search:               "machine learning",
	}

	_ = parsecore.Source{
		SourceName:           "BookRix",
		UrlREGEX:             "<big class=.item-title.><a class=.word-break. href=.([^\"']*)",
		IdREGEX:              "",
		DownloadUrlREGEX:     "data-download=.([^\"']*)",
		DownloadUrlComplete:  "https://www.bookrix.com/",
		TitleREGEX:           "<h2 class=.break-word.>([^<]*)",
		IsbnREGEX:            "",
		YearREGEX:            "",
		PublisherREGEX:       "",
		AuthorREGEX:          "<a rel=.author. href=[^>]*>([^<]*)",
		ExtensionREGEX:       "",
		PageREGEX:            "",
		LanguageREGEX:        "",
		SizeREGEX:            "",
		TimeREGEX:            "Publication Date:.([^<]*)",
		CompletePageUrlStart: "https://www.bookrix.com/books;page:",
		CompletePageUrlEnd:   ".html",
		IncompleteArticleUrl: "https://www.bookrix.com",
		AllUrls:              nil,
		Search:               "BookRixAll",
	}

	_ = parsecore.Source{
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
		CompletePageUrlStart: "?language=All&sort_by=field_downloads&page=",
		CompletePageUrlEnd:   "",
		AllUrls:              nil,
		Search:               "urlCat",
	}
}
