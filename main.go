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

}


func search(query string) {
	fmt.Printf("query: %s\n", query)
	site := "comick.io"
	books := scrapers.Registry[site].Search(query)
	for i, b := range books {
		fmt.Printf("\033[35m%d: \033[33m%s\033[32m/\033[94m%s\033[32m:\033[96m%s\n", i, site, b.Id, b.Title)
	}

}
