package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for i, link := range os.Args {
		if i == 0 {
			continue
		}
		for _, l := range process(link) {
			println(l)
		}
	}
}

func process(link string) []string {
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	links := getLinks(resp.Body)
	url, err := url.Parse(link)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(links); i++ {
		if !(strings.HasPrefix(links[i], "http:/") || strings.HasPrefix(links[i], "https:/")) {
			links[i] = path.Join(url.Host, url.Path, links[i])
		}
	}
	return links
}

func getLinks(fromReader io.ReadCloser) []string {
	defer fromReader.Close()
	var results []string
	tkns := html.NewTokenizer(fromReader)
	for {
		tknType := tkns.Next()
		switch {
		case tknType == html.ErrorToken:
			return results
		case tknType == html.StartTagToken:
			tkn := tkns.Token()
			if tkn.Data == "a" {
				for _, attr := range tkn.Attr {
					if attr.Key == "href" {
						results = append(results, attr.Val)
					}
				}
			}
		}
	}
}
