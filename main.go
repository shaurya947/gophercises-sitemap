package main

import (
	"flag"
	"fmt"
	"gophercises-sitemap/sitemap"
	"log"
)

func main() {
	urlStrPtr := flag.String("url", "https://www.petsmart.com/", "Site to build map of")
	flag.Parse()

	urls, err := sitemap.BuildFromUrlStr(*urlStrPtr)
	if err != nil {
		log.Fatalf("Error building sitemap for %s: %+v", *urlStrPtr, err)
	}

	for _, url := range urls {
		fmt.Println(url.String())
	}
}
