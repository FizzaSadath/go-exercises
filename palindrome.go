package main

// Write a function that accepts a string s as input and return
//       true if the string is a palindrome. false otherwise. A palindrome
//       is a string that reads the same from left to right and right to
//       left (e.g. "dad", "malayalam" etc.)
import (
	"fmt"
)

func palindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
func main() {
	s := "malayalam"
	res := palindrome(s)
	if res == true {
		fmt.Printf("%s is a Palindrome\n", s)
	} else {
		fmt.Printf("%s is not a Palindrome\n", s)
	}
}
