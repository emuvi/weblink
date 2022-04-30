package main

import (
	"io"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://www.uol.com.br/")
	if err != nil {
		panic(err)
	}
	links := parse(resp.Body)
	for _, link := range links {
		println(link)
	}
}

func parse(reader io.ReadCloser) (links []string) {
	defer reader.Close()
	var result []string
	tkns := html.NewTokenizer(reader)
	for {
		tknType := tkns.Next()
		switch {
		case tknType == html.ErrorToken:
			return result
		case tknType == html.StartTagToken:
			tkn := tkns.Token()
			if tkn.Data == "a" {
				for _, attr := range tkn.Attr {
					if attr.Key == "href" {
						result = append(result, attr.Val)
					}
				}
			}
		}
	}
}
