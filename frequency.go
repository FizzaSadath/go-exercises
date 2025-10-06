package main

// Write a function which accepts a string s as input and returns a
//       map whose keys are the characters in the string and whose values
//       are the number of time the character occurs. e.g. If "bobby" is
//       given as input, it will return a map with these values {b:3,
//       o:1, y:1}
import (
	"fmt"
)

func frequency(s string) map[string]int {
	m := map[string]int{}
	for i := 0; i < len(s); i++ {
		_, ok := m[string(s[i])]
		if ok {
			m[string(s[i])] += 1
		} else {
			m[string(s[i])] = 1
		}
	}
	return m

}
func main() {
	str := "malayalam"
	result := frequency(str)
	fmt.Println(result)
}
