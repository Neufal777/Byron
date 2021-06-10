package main

import (
	"flag"
	"fmt"

	"github.com/Byron/sources"
	"github.com/ttacon/chalk"
)

func main() {

	category := flag.String("cat", "math", "category")
	flag.Parse()
	fmt.Println(chalk.Magenta.Color("processing " + *category))

	//data.FilesOrganizer("Inventory/")

	cat1 := []string{"the", "of", "and", "a", "to", "in", "is", "you", "that", "it", "he", "was", "for", "on", "are", "six", "size", "skill", "skin", "small", "smile", "so", "social", "society", "soldier", "some", "somebody", "someone", "something", "sometimes", "son", "song", "soon", "sort", "sound", "source", "south", "southern", "space", "speak", "special", "specific", "speech", "spend", "sport", "spring", "staff", "stage", "stand", "standard", "star", "start", "state", "statement", "station", "stay", "step", "still", "stock", "stop", "store", "story", "strategy", "street", "strong", "structure", "student", "study", "stuff", "style", "subject", "success", "successful", "such", "suddenly", "suffer", "suggest", "summer", "support", "sure", "surface", "system", "table", "take", "talk", "task", "tax", "teach", "teacher", "team", "technology", "television", "tell", "ten", "tend", "term", "test", "than", "thank", "that", "the", "their", "them", "themselves", "then", "theory", "there", "these", "they", "thing", "think", "third", "this", "those", "though", "thought", "thousand", "threat", "three", "through", "throughout", "throw", "thus", "time", "to", "today", "together", "tonight", "too", "top", "total", "tough", "toward", "town", "trade", "traditional", "training", "travel", "treat", "treatment", "tree", "trial", "trip", "trouble", "true", "truth", "try", "turn", "TV", "two", "type", "under", "understand", "unit", "until", "up", "upon", "us", "use", "usually", "value", "various", "very", "victim", "view", "violence", "visit", "voice", "vote", "wait", "walk", "wall", "want", "war", "watch", "water", "way", "we", "weapon", "wear", "week", "weight", "well", "west", "western", "what"}
	cat2 := []string{"his", "they", "I", "at", "be", "this", "have", "from", "or", "one", "had", "by", "word", "but", "not", "relate", "relationship", "religious", "remain", "remember", "remove", "report", "represent", "Republican", "require", "research", "resource", "respond", "response", "responsibility", "rest", "result", "return", "reveal", "rich", "right", "rise", "risk", "road", "rock", "role", "room", "rule", "run", "safe", "same", "save", "say", "scene", "school", "science", "scientist", "score", "sea", "season", "seat", "second", "section", "security", "see", "seek", "seem", "sell", "send", "senior", "sense", "series", "serious", "serve", "service", "set", "seven", "several", "sex", "sexual", "shake", "share", "she", "shoot", "short", "shot", "should", "shoulder", "show", "side", "sign", "significant", "similar", "simple", "simply", "patient", "pattern", "pay", "peace", "people", "per", "perform", "performance", "perhaps", "period", "person", "personal", "phone", "physical", "pick", "picture", "piece", "place", "plan", "plant", "play", "player", "PM", "point", "police", "policy", "political", "politics", "poor", "popular", "population", "position", "positive", "possible", "power", "practice", "prepare", "present"}
	cat3 := []string{"as", "with", "what", "all", "were", "we", "your", "can", "said", "there", "use", "an", "air", "all", "allow", "almost", "alone", "along", "already", "also", "although", "always", "American", "among", "amount", "analysis", "and", "animal", "another", "answer", "any", "anyone", "anything", "appear", "apply", "approach", "area", "argue", "arm", "around", "arrive", "art", "article", "artist", "as", "ask", "assume", "at", "attack", "attention", "attorney", "audience", "author", "authority", "available", "avoid", "away", "baby", "back", "bad", "bag", "ball", "bank", "bar", "base", "be", "beat", "beautiful", "because", "become", "bed", "before", "begin", "behavior", "behind", "believe", "benefit", "best", "better", "between", "beyond", "big", "bill", "billion", "bit", "black", "blood", "blue", "radio", "raise", "range", "rate", "rather", "reach", "read", "ready", "real", "reality", "realize", "really", "reason", "receive", "recent", "recently", "recognize", "record", "red", "reduce", "reflect", "region"}
	cat4 := []string{"each", "which", "she", "do", "how", "their", "if", "will", "up", "other", "about", "out", "many", "board", "body", "book", "born", "both", "box", "boy", "break", "bring", "brother", "budget", "build", "building", "business", "but", "buy", "by", "call", "camera", "campaign", "can", "cancer", "candidate", "capital", "car", "card", "care", "career", "carry", "case", "catch", "cause", "cell", "center", "central", "century", "certain", "certainly", "chair", "challenge", "chance", "change", "character", "charge", "check", "child", "choice", "choose", "church", "citizen", "city", "civil", "claim", "class", "clear", "clearly", "close", "coach", "cold", "collection", "college", "color", "come", "commercial", "common", "community", "company", "compare", "computer", "concern", "condition", "conference", "Congress", "consider", "consumer", "contain", "continue", "control", "cost", "could", "since", "sing", "single", "sister", "sit", "site", "situation", "president", "pressure", "pretty", "prevent", "price", "private", "probably", "problem", "process", "produce", "product", "production", "professional", "professor", "program", "project", "property", "protect", "prove", "provide", "public", "pull", "purpose", "push", "put", "quality", "question", "quickly", "quite", "race"}
	cat5 := []string{"then", "them", "these", "so", "some", "her", "would", "make", "like", "him", "into", "time", "has", "look", "country", "couple", "course", "court", "cover", "create", "crime", "cultural", "culture", "cup", "current", "customer", "cut", "dark", "data", "daughter", "day", "dead", "deal", "death", "debate", "decade", "decide", "decision", "deep", "defense", "degree", "Democrat", "democratic", "describe", "design", "despite", "detail", "determine", "develop", "development", "die", "difference", "different", "difficult", "dinner", "direction", "director", "discover", "discuss", "discussion", "disease", "do", "doctor", "dog", "door", "down", "draw", "dream", "drive", "drop", "drug", "during", "each", "early", "east", "easy", "eat", "economic", "economy", "edge", "education", "effect", "effort", "eight", "either", "election", "else", "employee", "end", "energy", "enjoy", "enough", "enter", "entire", "environment", "environmental", "especially", "establish", "even", "evening", "event", "ever", "every", "everybody", "everyone", "everything", "evidence", "exactly", "example", "executive", "exist", "expect", "experience", "expert", "explain", "eye", "face", "fact", "factor", "fail", "fall", "family", "far", "fast", "father", "fear", "federal", "feel", "feeling", "few", "field", "fight", "figure", "fill", "film", "final", "finally", "financial", "find", "fine", "finger", "finish", "fire", "firm", "first", "fish", "five", "floor", "fly", "focus", "follow", "food", "foot", "for", "force", "foreign", "forget", "form", "former", "forward", "four", "free", "friend", "from", "front", "full", "fund", "future", "game", "garden", "gas", "general", "generation", "get", "girl", "give", "glass", "go", "goal", "good", "government", "great", "green", "ground", "group", "grow", "growth", "guess", "gun", "guy", "hair", "half", "hand", "hang", "happen", "happy", "hard", "have", "he", "head", "health", "hear", "heart", "heat", "heavy", "help", "her", "here", "herself", "high", "him", "himself", "his", "history", "hit", "hold", "home", "hope", "hospital", "hot", "hotel", "hour", "house", "how", "however", "huge", "human", "hundred", "husband", "I", "idea", "identify", "if", "image", "imagine", "impact", "important", "improve", "in", "include", "including", "increase", "indeed", "indicate", "individual", "industry", "information", "inside", "instead", "institution", "interest", "interesting", "international", "interview", "into", "investment", "involve", "issue", "it", "item", "its", "itself", "job", "join", "just", "keep", "key", "kid", "kill", "kind", "kitchen", "know", "knowledge", "land", "language", "large", "last", "late", "later", "laugh", "law", "lawyer", "lay", "lead", "leader", "learn", "least", "leave", "left", "leg", "legal", "less", "let", "letter", "level", "lie", "life", "light", "like", "likely", "line", "list", "listen", "little", "live", "local", "long", "look", "lose", "loss", "lot", "love", "low", "machine", "magazine", "main", "maintain", "major", "majority", "make", "man", "manage", "management", "manager", "many", "market", "marriage", "material", "matter", "may", "maybe", "me", "mean", "measure", "media", "medical", "meet", "meeting", "member", "memory", "mention", "message", "method", "middle", "might", "military", "million", "mind", "minute", "miss", "mission", "model", "modern", "moment", "money", "month", "more", "morning", "most", "mother", "mouth", "move", "movement", "movie", "Mr", "Mrs", "much", "music", "must", "my", "myself", "name", "nation", "national", "natural", "nature", "near", "nearly", "necessary", "need", "network", "never", "new", "news", "newspaper", "next", "nice", "night", "no", "none", "nor", "north", "not", "note", "nothing", "notice", "now", "n't", "number", "occur", "of", "off", "offer", "office", "officer", "official", "often", "oh", "oil", "ok", "old", "on", "once", "one", "only", "onto", "open", "operation", "opportunity", "option", "or", "order", "organization", "other", "others", "our", "out", "outside", "over", "own", "owner", "page", "pain", "painting", "paper", "parent", "part", "participant", "particular", "particularly", "partner", "party", "pass", "past"}
	cat6 := []string{"two", "more", "write", "go", "see", "number", "no", "way", "could", "people", "my", "than", "first", "water", "write", "writer", "wrong", "yard", "yeah", "year", "yes", "yet", "you", "young", "your", "yourself"}
	cat7 := []string{"been", "call", "oil", "its", "now", "find", "long", "down", "day", "did", "get", "come", "made", "may", "part"}
	cat8 := []string{"ability", "able", "about", "above", "accept", "according", "account", "across", "act", "action", "activity", "actually", "add", "address", "administration", "admit", "adult", "affect", "after", "again", "against", "age", "agency", "agent", "ago", "agree", "agreement", "ahead"}
	cat9 := []string{"whatever", "when", "where", "whether", "which", "while", "white", "who", "whole", "whom", "whose", "why", "wide", "wife", "will", "win", "wind", "window", "wish", "with", "within", "without", "woman", "wonder", "word", "work", "worker", "world", "worry", "would"}
	science := []string{"machine learning", "artificial Inteligence", "neuroscience", "computer", "hack", "maths", "biology", "medicine", "mit", "genomic", "physics", "fisica", "chemistry", "universe", "paper", "science", "machine"}

	switch *category {
	case "1":
		for _, c := range cat1 {
			source := sources.Source{
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
				CompletePageUrl:      "https://libgen.is/search.php?&res=100&req=" + c + "&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
				IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
				AllUrls:              nil,
				Search:               c,
			}
			fmt.Println(chalk.Magenta.Color("processing " + c))
			source.GetArticles()
		}
	case "2":
		for _, c := range cat2 {
			source := sources.Source{
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
				CompletePageUrl:      "https://libgen.is/search.php?&res=100&req=" + c + "&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
				IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
				AllUrls:              nil,
				Search:               c,
			}
			fmt.Println(chalk.Magenta.Color("processing " + c))
			source.GetArticles()
		}
	case "3":
		for _, c := range cat3 {
			source := sources.Source{
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
				CompletePageUrl:      "https://libgen.is/search.php?&res=100&req=" + c + "&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
				IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
				AllUrls:              nil,
				Search:               c,
			}
			fmt.Println(chalk.Magenta.Color("processing " + c))
			source.GetArticles()
		}
	case "4":
		for _, c := range cat4 {
			source := sources.Source{
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
				CompletePageUrl:      "https://libgen.is/search.php?&res=100&req=" + c + "&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
				IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
				AllUrls:              nil,
				Search:               c,
			}
			fmt.Println(chalk.Magenta.Color("processing " + c))
			source.GetArticles()
		}
	case "5":
		for _, c := range cat5 {
			source := sources.Source{
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
				CompletePageUrl:      "https://libgen.is/search.php?&res=100&req=" + c + "&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
				IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
				AllUrls:              nil,
				Search:               c,
			}
			fmt.Println(chalk.Magenta.Color("processing " + c))
			source.GetArticles()
		}
	case "6":
		for _, c := range cat6 {
			source := sources.Source{
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
				CompletePageUrl:      "https://libgen.is/search.php?&res=100&req=" + c + "&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
				IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
				AllUrls:              nil,
				Search:               c,
			}
			fmt.Println(chalk.Magenta.Color("processing " + c))
			source.GetArticles()
		}
	case "7":
		for _, c := range cat7 {
			source := sources.Source{
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
				CompletePageUrl:      "https://libgen.is/search.php?&res=100&req=" + c + "&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
				IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
				AllUrls:              nil,
				Search:               c,
			}
			fmt.Println(chalk.Magenta.Color("processing " + c))
			source.GetArticles()
		}
	case "8":
		for _, c := range cat8 {
			source := sources.Source{
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
				CompletePageUrl:      "https://libgen.is/search.php?&res=100&req=" + c + "&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
				IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
				AllUrls:              nil,
				Search:               c,
			}
			fmt.Println(chalk.Magenta.Color("processing " + c))
			source.GetArticles()
		}
	case "9":
		for _, c := range cat9 {
			source := sources.Source{
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
				CompletePageUrl:      "https://libgen.is/search.php?&res=100&req=" + c + "&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
				IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
				AllUrls:              nil,
				Search:               c,
			}
			fmt.Println(chalk.Magenta.Color("processing " + c))
			source.GetArticles()
		}
	default:
		for _, c := range science {
			source := sources.Source{
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
				CompletePageUrl:      "https://libgen.is/search.php?&res=100&req=" + c + "&phrase=1&view=simple&column=def&sort=year&sortmode=DESC&page=",
				IncompleteArticleUrl: "https://libgen.is/book/index.php?md5=",
				AllUrls:              nil,
				Search:               c,
			}
			fmt.Println(chalk.Magenta.Color("processing " + c))
			source.GetArticles()
		}
	}
}
