package lib

import (
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"golang.org/x/net/html"
)

type Filters struct {
	FollowExternal bool
	OnlyUppers     bool
	OnlyLowers     bool
}

func Parse(root string, filters Filters) []string {
	resp, err := http.Get(root)
	if err != nil {
		panic(err)
	}
	links := Catch(resp.Body)
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
		if Filter(urlRoot, urlLink, filters) {
			results = append(results, urlLink.String())
		}
	}
	return results
}

func Filter(urlRoot *url.URL, urlLink *url.URL, filters Filters) bool {
	if !filters.FollowExternal && IsExternal(urlRoot, urlLink) {
		return false
	}
	if filters.OnlyUppers && !IsSameOrUpper(urlRoot, urlLink) {
		return false
	}
	if filters.OnlyLowers && !IsSameOrLower(urlRoot, urlLink) {
		return false
	}
	return true
}

func IsExternal(urlRoot *url.URL, urlLink *url.URL) bool {
	return urlLink.Host != urlRoot.Host
}

func IsSameOrUpper(urlRoot *url.URL, urlLink *url.URL) bool {
	if IsExternal(urlLink, urlRoot) {
		return false
	}
	dirRoot, _ := path.Split(urlRoot.Path)
	dirLink, _ := path.Split(urlLink.Path)
	return strings.HasPrefix(dirRoot, dirLink)
}

func IsSameOrLower(urlRoot *url.URL, urlLink *url.URL) bool {
	if IsExternal(urlLink, urlRoot) {
		return false
	}
	dirRoot, _ := path.Split(urlRoot.Path)
	dirLink, _ := path.Split(urlLink.Path)
	return strings.HasPrefix(dirLink, dirRoot)
}

func Catch(fromReader io.ReadCloser) []string {
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
