package main

import (
	"flag"
	"fmt"
	"strings"

	"tuiteka/scrapers"
	"tuiteka/reader"
)


func main() {

	flag.Parse()
	args := flag.Args()
	fmt.Println(len(args))



	fmt.Println("Hello there!")

	if len(args) < 2 {
		fmt.Println("nothin")
		return
	}


	if args[0] == "i" || args[0] == "interactive" {
		fmt.Print("interactive\n")
		query := strings.Join(args[1:], " ")
		interactive(query)
		return
	}

	if args[0] == "s" || args[0] == "search" {
		query := strings.Join(args[1:], " ")
		search(query)
		return
	}

	if args[0] == "c" || args[0] == "chapters" {
		id := strings.Join(args[1:], " ")
		chapters(id)
		return
	}

	if args[0] == "r" || args[0] == "read" {

		if len(args) != 3 {
			fmt.Println("not enough arguments")
			return
		}

		slug := args[1]
		chapter := args[2]
		pages(slug, chapter)
		return
	}

}


func search(query string) {
	fmt.Printf("query: %s\n", query)
	site := "comick.io"
	books := scrapers.Registry[site].Search(query)
	for i, b := range books {
		fmt.Printf("\033[35m%d: \033[33m%s\033[32m/\033[94m%s\033[32m:\033[96m%s\n", i, site, b.Id, b.Title)
	}

}

func chapters(id string) {
	fmt.Printf("id: %s\n", id)
	site := "comick.io"
	results := scrapers.Registry[site].Chapters(id)
	for i, c := range results {
		fmt.Printf("\033[35m%d: \033[33m%s\033[32m/\033[94m%s\033[32m:\033[96m%s \033[34m{ %s }\n", i, site, c.Id, c.Number, c.Info)
	}

}

func pages(slug string, chapter string) {
	fmt.Printf("read: %s/%s\n", slug, chapter)
	site := "comick.io"
	scrapers.Registry[site].Pages(slug, chapter)
	// results := scrapers.Registry[site].Pages(id, chapter)
	// for i, c := range results {
	// 	fmt.Printf("\033[35m%d: \033[33m%s\033[32m/\033[94m%s\033[32m:\033[96m%s \033[34m{ %s }\n", i, site, c.Id, c.Number, c.Info)
	// }

}

func interactive(query string) {
	site := "comick.io"
	images := scrapers.Registry[site].Interactive(query)


	reader.Read(images)

}
