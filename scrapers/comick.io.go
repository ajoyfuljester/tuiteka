package scrapers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func search(query string) []Book {

	baseURL := "https://api.comick.io/v1.0/search"
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	formattedQuery := query

	q := u.Query()
	q.Set("q", formattedQuery)
	// i don't know what exactly it is, but if set to false, response has additional information
	q.Set("t", "true")

	u.RawQuery = q.Encode()


	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic(err)
	}

	headers := BaseHeaders.Clone()
	headers.Set("Referer", "https://comick.io/")
	headers.Set("Origin", "https://comick.io")

	for k, h := range headers {
		fmt.Printf("%s: %s\n", k, h)
	}

	req.Header = headers


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}


	type Data struct {
		Id string `json:"hid"`
		Title string `json:"title"`
		Type string `json:"type"`
	}

	var data []Data
	fmt.Print(string(body))
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	books := make([]Book, 0, len(data))

	for _, d := range data {
		if d.Type == "author" {
			continue
		}
		books = append(books, Book{
			Id: d.Id,
			Title: d.Title,
		})

	}

	return books

	// curl 'https://api.comick.io/v1.0/search?q=angel+nex^&t=true' \
	// --compressed \
}

func init() {
	Register(
		"comick.io",
		ScraperFuncs{
			Search: search,
		},
	)
}

