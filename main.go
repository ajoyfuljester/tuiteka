package main

import (
	"flag"
	"fmt"
	"strings"

	"tuiteka/scrapers"
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
		fmt.Printf("\033[35m%d: \033[33m%s\033[32m/\033[94m%s\033[32m:\033[96m%s \033[34m{ %s }\n", i, site, c.Id, c.Title, c.Info)
	}

}
