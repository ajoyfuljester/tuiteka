package scrapers

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"net/http"
	"net/url"
	"strings"

	"tuiteka/utils"
	"github.com/PuerkitoBio/goquery"
)


var site = "comick.io"


func init() {

	getHeaders := func() http.Header {
		headers := BaseHeaders.Clone()
		headers.Set("Referer", "https://comick.io/")
		headers.Set("Origin", "https://comick.io")

		return headers

	}

	search := func(query string) []Book {
		headers := getHeaders()

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
			Slug string `json:"slug"`
		}

		var data []Data
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
				Slug: d.Slug,
			})

		}

		return books
	}

	// TODO: probably pages of these requests

	getChapters := func(id string) []Chapter {
		headers := getHeaders()

		requestURL := fmt.Sprintf("https://api.comick.io/comic/%s/chapters", id)

		req, err := http.NewRequest("GET", requestURL, nil)
		if err != nil {
			panic(err)
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

		// TODO: handle if id is wrong


		type ChapterWrapper struct {
			Id string `json:"hid"`
			Number string `json:"chap"`
			Language string `json:"lang"`
			Groups []string `json:"group_name"`
		}

		type Data struct {
			Chapters []ChapterWrapper `json:"chapters"`
		}

		var data Data
		err = json.Unmarshal(body, &data)
		if err != nil {
			panic(err)
		}

		chaps := make([]Chapter, 0, len(data.Chapters))

		for _, c := range data.Chapters {
			chaps = append(chaps, Chapter{
				Id: c.Id,
				Number: c.Number,
				Language: c.Language,
				Info: fmt.Sprintf("[%s]", strings.Join(c.Groups, ", ")),
			})

		}

		return chaps

	}



	getPages := func(book Book, chapter Chapter) []image.Image {
		headers := getHeaders()

		slug := book.Slug

		baseURL := fmt.Sprintf("https://comick.io/comic/%s/%s-chapter-%s-%s", slug, chapter.Id, chapter.Number, chapter.Language)
		fmt.Printf("URL: %s\n", baseURL)
		u, err := url.Parse(baseURL)
		if err != nil {
			panic(err)
		}


		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			panic(err)
		}

		req.Header = headers


		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()


		document, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			panic(err)
		}


		selection := document.Find("#__NEXT_DATA__")
		dataString := selection.Text()

		type PageData struct {
			FileName string `json:"b2key"`
		}

		var dataWrapper struct {
			Props struct {
				PageProps struct {
					Chapter struct {
						Data []PageData `json:"md_images"`
					} `json:"chapter"`
				} `json:"pageProps"`
			} `json:"props"`

		}

		err = json.Unmarshal([]byte(dataString), &dataWrapper)

		data := dataWrapper.Props.PageProps.Chapter.Data

		for i, d := range data {
			fmt.Printf("\033[37m%d: %s\n", i, d.FileName)
		}


		urls := utils.Map(data, func(pd PageData) string {
			return fmt.Sprintf("https://meo.comick.pictures/%s", pd.FileName)
		})



		images, err := getImages(urls, headers)
		if err != nil {
			panic(err)
		}


		return images
	}

	interactive := func(query string) []image.Image {
		books := search(query)


		for i, b := range books {
			fmt.Printf("\033[35m%d: \033[33m%s\033[32m/\033[94m%s\033[32m:\033[96m%s\n", i, site, b.Id, b.Title)
		}

		var index int16
		fmt.Print("\033[35mChoose book number: ")
		_, err := fmt.Scan(&index)
		if err != nil {
			fmt.Printf("Out of range %d:%d\n", index, len(books))
			panic(err)
		}
		fmt.Printf("Chosen %d\n", index)

		chosenBook := books[index]
		chapters := getChapters(chosenBook.Id)


		for i, c := range chapters {
			fmt.Printf("\033[35m%d: \033[33m%s\033[32m/\033[94m%s\033[32m:\033[96m%s \033[34m{ %s }\n", i, site, c.Id, c.Number, c.Info)
		}

		fmt.Print("\033[35mChoose chapter number: ")
		_, err = fmt.Scan(&index)
		if err != nil {
			fmt.Printf("Out of range %d:%d\n", index, len(chapters))
			panic(err)
		}

		fmt.Printf("Chosen %d\n", index)
		chosenChapter := chapters[index]

		images := getPages(chosenBook, chosenChapter)


		return images


	}






	Register(
		site,
		ScraperFuncs{
			Search: search,
			Chapters: getChapters,
			Interactive: interactive,
		},
	)
}

