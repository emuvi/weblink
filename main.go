package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

var until_depth = 3
var follow_external = false

func main() {
	links := make(map[string]int)
	index := 1
	length := len(os.Args)
	var err error
	for index < length {
		if os.Args[index] == "-d" || os.Args[index] == "--depth" {
			until_depth, err = strconv.Atoi(os.Args[index+1])
			if err != nil {
				panic(err)
			}
			index += 2
		} else {
			links[strings.ToLower(os.Args[index])] = 1
			index++
		}
	}
	actual_depth := 1
	for actual_depth <= until_depth {
		for link, link_depth := range links {
			if link_depth == actual_depth {
				for _, new_link := range parse(link) {
					links[new_link] = actual_depth + 1
				}
			}
		}
		actual_depth += 1
	}
	for link := range links {
		println(link)
	}
}

func parse(root string) []string {
	resp, err := http.Get(root)
	if err != nil {
		panic(err)
	}
	links := catch(resp.Body)
	urlRoot, err := url.Parse(root)
	if err != nil {
		panic(err)
	}
	var results []string
	for _, link := range links {
		urlLink, err := url.Parse(link)
		if err != nil {
			panic(err)
		}
		if !urlLink.IsAbs() {
			urlLink = urlRoot.ResolveReference(urlLink)
		}

		if strings.HasPrefix(link, root) {
			results = append(results, link)
		}
	}
	return results
}

func catch(fromReader io.ReadCloser) []string {
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
						results = append(results, strings.ToLower(attr.Val))
					}
				}
			}
		}
	}
}
