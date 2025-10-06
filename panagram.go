package main

import (
	"fmt"
	"strings"
)

func panagram(s string) bool {
	s = strings.ToLower(s)
	for ch := 'a'; ch <= 'z'; ch++ {
		if !strings.Contains(s, string(ch)) {
			return false
		}
	}
	return true
}
func main() {
	s := "The quick brown fox jumps over the lazy dog"
	if panagram(s) == true {
		fmt.Printf(" \"%s\" is a panagram\n", s)
	} else {
		fmt.Printf(" \"%s\" is not a panagram\n", s)
	}
}
