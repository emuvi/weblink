package main

import (
	"os"
	"strconv"
	"strings"
	"weblinks/lib"
)

func main() {
	var until_depth = 1
	var follow_external = false
	var only_uppers = false
	var only_lowers = false
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
			index++
		} else if os.Args[index] == "-e" || os.Args[index] == "--external" {
			follow_external = true
		} else if os.Args[index] == "-u" || os.Args[index] == "--only-uppers" {
			only_uppers = true
		} else if os.Args[index] == "-l" || os.Args[index] == "--only-lowers" {
			only_lowers = true
		} else {
			links[strings.ToLower(os.Args[index])] = 1
		}
		index++
	}
	actual_depth := 1
	for actual_depth <= until_depth {
		for link, link_depth := range links {
			if link_depth == actual_depth {
				for _, new_link := range lib.Parse(link, lib.Filters{
					FollowExternal: follow_external,
					OnlyUppers:     only_uppers,
					OnlyLowers:     only_lowers,
				}) {
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
