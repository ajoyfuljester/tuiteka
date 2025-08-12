package scrapers

type Book struct {
	Id string
	Title string
}

type Chapter struct {
	Id string
	Title string
	Info string
}

type ScraperFuncs struct {
	Search func(string) []Book
	Chapters func(string) []Chapter
}

var Registry = make(map[string]ScraperFuncs)

func Register(name string, fns ScraperFuncs) {
	Registry[name] = fns
}
