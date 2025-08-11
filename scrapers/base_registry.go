package scrapers

type Book struct {
	Id string
	Title string
}

type ScraperFuncs struct {
	Search func(string) []Book
}

var Registry = make(map[string]ScraperFuncs)

func Register(name string, fns ScraperFuncs) {
	Registry[name] = fns
}
