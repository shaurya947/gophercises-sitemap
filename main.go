package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"gophercises-sitemap/sitemap"
	"log"
)

type Url struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	Urls    []Url
}

func main() {
	urlStrPtr := flag.String("url", "https://www.petsmart.com/", "Site to build map of")
	flag.Parse()

	rawUrls, err := sitemap.BuildFromUrlStr(*urlStrPtr)
	if err != nil {
		log.Fatalf("Error building sitemap for %s: %+v", *urlStrPtr, err)
	}

	urls := []Url{}
	for _, rawUrl := range rawUrls {
		urls = append(urls, Url{Loc: rawUrl.String()})
	}

	urlSet := &UrlSet{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Urls:  urls,
	}

	rawXml, err := xml.MarshalIndent(urlSet, "", "  ")
	if err != nil {
		log.Fatalln("Error marshaling XML")
	}
	fmt.Println(xml.Header + string(rawXml))
}
