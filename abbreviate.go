package main

// 			Write a function which takes a string s as input and returns the
//       short form of s by taking the first characters of all the words
//       which are in caps. e.g. if the input is "Indian Institute of
//       Management", it will return "IIM" (the "of" is skipped because
//       it doesn't start with a capital letter)

import (
	"fmt"
	"strings"
)

func main() {
	var s, sf string
	s = "Indian Institue of Management"
	sf = abbreviate(s)
	fmt.Println(sf)
}
func abbreviate(s1 string) string {
	var letters []string
	words := strings.Fields(s1)
	for _, word := range words {
		if word[0] >= 'A' && word[0] <= 'Z' {
			letters = append(letters, string(word[0]))
		}

	}
	shortform := strings.Join(letters, "")
	return shortform
}
